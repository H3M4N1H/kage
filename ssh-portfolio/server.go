package main

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"
	"errors"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/muesli/termenv"
	gossh "golang.org/x/crypto/ssh"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"
)

type config struct {
	Address, HostKey          string
	MaxConnections            int
	Handshake, Idle, Duration time.Duration
}

func defaultConfig() config {
	return config{"127.0.0.1:2222", ".state/host_ed25519", 32, 10 * time.Second, 2 * time.Minute, 10 * time.Minute}
}
func hostKey(path string) (gossh.Signer, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return nil, err
	}
	info, err := os.Lstat(path)
	if errors.Is(err, os.ErrNotExist) {
		_, key, e := ed25519.GenerateKey(rand.Reader)
		if e != nil {
			return nil, e
		}
		block, e := gossh.MarshalPrivateKey(key, "")
		if e != nil {
			return nil, e
		}
		f, e := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
		if e != nil {
			return nil, e
		}
		_, e = f.Write(pem.EncodeToMemory(block))
		closeErr := f.Close()
		if e != nil {
			return nil, e
		}
		if closeErr != nil {
			return nil, closeErr
		}
	} else if err != nil {
		return nil, err
	} else if !info.Mode().IsRegular() || info.Mode().Perm()&0077 != 0 {
		return nil, fmt.Errorf("host key must be a regular file with permissions 0600 or stricter")
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return gossh.ParsePrivateKey(b)
}

type connectionState struct {
	conn *guardedConn
	used atomic.Bool
}
type connKey struct{}

// Filter requests before the library session handler. No agent forwarding,
// environment accumulation, oversized terminals, or arbitrary request types.
type filteredChannel struct {
	gossh.NewChannel
	ctx context.Context
}

func (n filteredChannel) Accept() (gossh.Channel, <-chan *gossh.Request, error) {
	ch, in, err := n.NewChannel.Accept()
	if err != nil {
		return nil, nil, err
	}
	out := make(chan *gossh.Request)
	go func() {
		defer close(out)
		count := 0
		for {
			select {
			case <-n.ctx.Done():
				return
			case r, ok := <-in:
				if !ok {
					return
				}
				count++
				if count > 2048 || len(r.Payload) > 4096 {
					_ = ch.Close()
					return
				}
				allowed := r.Type == "shell"
				if r.Type == "pty-req" {
					var p struct {
						Term                                   string
						Width, Height, PixelWidth, PixelHeight uint32
						Modes                                  string
					}
					allowed = gossh.Unmarshal(r.Payload, &p) == nil && len(p.Term) <= 64 && p.Width > 0 && p.Width <= 500 && p.Height > 0 && p.Height <= 200
				}
				if r.Type == "window-change" {
					var w struct{ Width, Height, PixelWidth, PixelHeight uint32 }
					allowed = gossh.Unmarshal(r.Payload, &w) == nil && w.Width > 0 && w.Width <= 500 && w.Height > 0 && w.Height <= 200
				}
				if !allowed {
					_ = r.Reply(false, nil)
					continue
				}
				select {
				case out <- r:
				case <-n.ctx.Done():
					return
				}
			}
		}
	}()
	return ch, out, nil
}
func makeServer(c config, content Content) (*ssh.Server, error) {
	signer, err := hostKey(c.HostKey)
	if err != nil {
		return nil, err
	}
	s, err := wish.NewServer(wish.WithAddress(c.Address))
	if err != nil {
		return nil, err
	}
	s.AddHostKey(signer)
	s.Version = "SSH-2.0-HemanthPortfolio"
	s.ConnCallback = func(ctx ssh.Context, c net.Conn) net.Conn {
		g, ok := c.(*guardedConn)
		if !ok {
			return nil
		}
		ctx.SetValue(connKey{}, &connectionState{conn: g})
		return c
	}
	s.ServerConfigCallback = func(ctx ssh.Context) *gossh.ServerConfig {
		return &gossh.ServerConfig{NoClientAuth: true, MaxAuthTries: 1}
	}
	s.RequestHandlers = map[string]ssh.RequestHandler{}
	s.SubsystemHandlers = map[string]ssh.SubsystemHandler{}
	s.ChannelHandlers = map[string]ssh.ChannelHandler{"session": func(srv *ssh.Server, c *gossh.ServerConn, ch gossh.NewChannel, ctx ssh.Context) {
		state := ctx.Value(connKey{}).(*connectionState)
		if !state.used.CompareAndSwap(false, true) {
			_ = ch.Reject(gossh.ResourceShortage, "one portfolio session per connection")
			return
		}
		state.conn.opened.Store(true)
		ssh.DefaultSessionHandler(srv, c, filteredChannel{ch, ctx}, ctx)
	}}
	s.SessionRequestCallback = func(sess ssh.Session, kind string) bool { return kind == "shell" && sess.RawCommand() == "" }
	s.Handler = func(sess ssh.Session) {
		pty, changes, ok := sess.Pty()
		if !ok || pty.Term == "dumb" {
			_, _ = fmt.Fprint(sess, plainContent(content))
			return
		}
		renderer := lipgloss.NewRenderer(sess)
		profile := termenv.Ascii
		if strings.Contains(pty.Term, "256color") {
			profile = termenv.ANSI256
		} else if strings.Contains(pty.Term, "xterm") || strings.Contains(pty.Term, "screen") || strings.Contains(pty.Term, "vt100") {
			profile = termenv.ANSI
		}
		if os.Getenv("NO_COLOR") != "" {
			profile = termenv.Ascii
		}
		renderer.SetColorProfile(profile)
		m := newModel(content, renderer.NewStyle().Foreground(lipgloss.Color("#B5F36B")))
		m.width = clamp(pty.Window.Width, 1, 160)
		m.height = clamp(pty.Window.Height, 1, 60)
		ctx, cancel := context.WithCancel(sess.Context())
		defer cancel()
		p := tea.NewProgram(m, tea.WithInput(sess), tea.WithOutput(sess), tea.WithAltScreen(), tea.WithContext(ctx), tea.WithFPS(20), tea.WithoutSignalHandler())
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case w, open := <-changes:
					if !open {
						return
					}
					p.Send(tea.WindowSizeMsg{Width: w.Width, Height: w.Height})
				}
			}
		}()
		_, _ = p.Run()
	}
	return s, nil
}
