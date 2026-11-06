package main

import (
	"bytes"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	gossh "golang.org/x/crypto/ssh"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestContent(t *testing.T) {
	c, e := loadContent()
	if e != nil || len(c.Sections) != 8 {
		t.Fatal(e)
	}
	for _, b := range []string{`{}`, `bad`, `{"name":"x","sections":[]}`} {
		if _, e := parseContent([]byte(b)); e == nil {
			t.Fatal("invalid content accepted")
		}
	}
}
func TestNavigationAndResize(t *testing.T) {
	c, _ := loadContent()
	m := newModel(c, lipgloss.NewStyle())
	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m = next.(model)
	if m.cursor != 1 {
		t.Fatal("down")
	}
	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = next.(model)
	if m.page != 1 {
		t.Fatal("enter")
	}
	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	m = next.(model)
	if m.page != -1 {
		t.Fatal("back")
	}
	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'9'}})
	if next.(model).page != -1 {
		t.Fatal("invalid selection")
	}
	for i := 1; i <= 8; i++ {
		next, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{rune('0' + i)}})
		if next.(model).page != i-1 {
			t.Fatal("number routing")
		}
	}
	for _, w := range []int{1, 20, 24, 40, 80, 160, 10000} {
		for _, h := range []int{1, 10, 24, 40, 10000} {
			next, _ = m.Update(tea.WindowSizeMsg{Width: w, Height: h})
			n := next.(model)
			for _, line := range strings.Split(n.View(), "\n") {
				if ansi.StringWidth(line) > min(w, 160) {
					t.Fatalf("overflow width %d: %q", w, line)
				}
			}
		}
	}
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	if cmd == nil {
		t.Fatal("quit missing")
	}
}
func TestWrap(t *testing.T) {
	for _, s := range []string{"longwordwithoutbreaks", strings.Repeat("hello ", 40), "你好世界 👋 testing", "https://www.linkedin.com/in/p-hemanth-gupta-657b05353"} {
		for _, line := range strings.Split(wrap(s, 12), "\n") {
			if ansi.StringWidth(line) > 12 {
				t.Fatal(line)
			}
		}
	}
}
func TestLimiter(t *testing.T) {
	l := newLimiter(3)
	now := time.Now()
	if !l.acquire("a", now) || !l.acquire("a", now) || l.acquire("a", now) {
		t.Fatal("IP cap")
	}
	if !l.acquire("b", now) || l.acquire("c", now) {
		t.Fatal("global cap")
	}
	l.release("a")
	l.release("a")
	for i := 0; i < 3; i++ {
		if !l.acquire("a", now) {
			t.Fatal("early rate limit")
		}
		l.release("a")
	}
	if l.acquire("a", now) {
		t.Fatal("rate limit missing")
	}
	if !l.acquire("a", now.Add(time.Minute)) {
		t.Fatal("rate not reset")
	}
}
func TestHostKey(t *testing.T) {
	path := filepath.Join(t.TempDir(), "key")
	a, e := hostKey(path)
	if e != nil {
		t.Fatal(e)
	}
	b, e := hostKey(path)
	if e != nil || !bytes.Equal(a.PublicKey().Marshal(), b.PublicKey().Marshal()) {
		t.Fatal("key not persistent")
	}
	_ = os.Chmod(path, 0644)
	if _, e = hostKey(path); e == nil {
		t.Fatal("unsafe permissions accepted")
	}
}
func startTestServer(t *testing.T, adjust func(*config)) (string, gossh.PublicKey) {
	t.Helper()
	c := defaultConfig()
	c.HostKey = filepath.Join(t.TempDir(), "key")
	if adjust != nil {
		adjust(&c)
	}
	content, _ := loadContent()
	s, e := makeServer(c, content)
	if e != nil {
		t.Fatal(e)
	}
	ln, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	ll := &limitedListener{ln, newLimiter(c.MaxConnections), c.Handshake, c.Idle, c.Duration}
	go func() { _ = s.Serve(ll) }()
	t.Cleanup(func() { _ = s.Close() })
	key, e := hostKey(c.HostKey)
	if e != nil {
		t.Fatal(e)
	}
	return ln.Addr().String(), key.PublicKey()
}
func client(t *testing.T, addr string, key gossh.PublicKey) *gossh.Client {
	t.Helper()
	c, e := gossh.Dial("tcp", addr, &gossh.ClientConfig{User: "visitor", HostKeyCallback: gossh.FixedHostKey(key), Timeout: 2 * time.Second})
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { _ = c.Close() })
	return c
}
func TestPlainSSHAndDeniedRequests(t *testing.T) {
	addr, key := startTestServer(t, nil)
	c := client(t, addr, key)
	session, e := c.NewSession()
	if e != nil {
		t.Fatal(e)
	}
	defer session.Close()
	for _, typ := range []string{"auth-agent-req@openssh.com", "x11-req", "env", "exec", "subsystem"} {
		ok, e := session.SendRequest(typ, true, nil)
		if e != nil || ok {
			t.Fatalf("%s not denied: %v %v", typ, ok, e)
		}
	}
	var out bytes.Buffer
	session.Stdout = &out
	if e = session.Shell(); e != nil {
		t.Fatal(e)
	}
	if e = session.Wait(); e != nil {
		t.Fatal(e)
	}
	if !strings.Contains(out.String(), "Hemanth") || strings.Contains(out.String(), "\x1b") {
		t.Fatal("plain output")
	}
	if _, e = c.NewSession(); e == nil {
		t.Fatal("second session allowed")
	}
	if _, _, e = c.OpenChannel("direct-tcpip", nil); e == nil {
		t.Fatal("TCP forwarding allowed")
	}
	if ok, _, e := c.SendRequest("tcpip-forward", true, nil); e != nil || ok {
		t.Fatal("remote forwarding allowed")
	}
}
func TestExecAndSFTP(t *testing.T) {
	for _, mode := range []string{"exec", "sftp"} {
		t.Run(mode, func(t *testing.T) {
			addr, key := startTestServer(t, nil)
			c := client(t, addr, key)
			s, e := c.NewSession()
			if e != nil {
				t.Fatal(e)
			}
			defer s.Close()
			if mode == "exec" {
				e = s.Run("cat /etc/passwd")
			} else {
				e = s.RequestSubsystem("sftp")
			}
			if e == nil {
				t.Fatal("forbidden request accepted")
			}
		})
	}
}
func TestInteractiveSSH(t *testing.T) {
	addr, key := startTestServer(t, nil)
	c := client(t, addr, key)
	s, e := c.NewSession()
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if e = s.RequestPty("xterm-256color", 24, 80, gossh.TerminalModes{}); e != nil {
		t.Fatal(e)
	}
	in, _ := s.StdinPipe()
	out, _ := s.StdoutPipe()
	if e = s.Shell(); e != nil {
		t.Fatal(e)
	}
	data := make(chan string, 1)
	go func() { b, _ := io.ReadAll(out); data <- string(b) }()
	time.Sleep(100 * time.Millisecond)
	_, _ = in.Write([]byte("3"))
	time.Sleep(100 * time.Millisecond)
	_ = s.WindowChange(16, 40)
	time.Sleep(100 * time.Millisecond)
	_, _ = in.Write([]byte("q"))
	select {
	case b := <-data:
		if !strings.Contains(b, "HEMANTH") || !strings.Contains(b, "Aegis") {
			t.Fatalf("missing UI: %q", b)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("interactive session did not exit")
	}
	if e = s.Wait(); e != nil {
		t.Fatal(e)
	}
}
func TestTimeouts(t *testing.T) {
	t.Run("handshake", func(t *testing.T) {
		addr, _ := startTestServer(t, func(c *config) { c.Handshake = 80 * time.Millisecond })
		c, e := net.Dial("tcp", addr)
		if e != nil {
			t.Fatal(e)
		}
		defer c.Close()
		_ = c.SetReadDeadline(time.Now().Add(time.Second))
		b, readErr := io.ReadAll(c)
		if readErr != nil {
			t.Fatalf("server did not close handshake before client deadline: %v", readErr)
		}
		if len(b) == 0 {
			t.Fatal("missing handshake banner")
		}
	})
	for _, which := range []string{"idle", "duration"} {
		t.Run(which, func(t *testing.T) {
			addr, key := startTestServer(t, func(c *config) {
				if which == "idle" {
					c.Idle = 120 * time.Millisecond
				} else {
					c.Duration = 120 * time.Millisecond
				}
			})
			c := client(t, addr, key)
			s, e := c.NewSession()
			if e != nil {
				t.Fatal(e)
			}
			defer s.Close()
			done := make(chan error, 1)
			go func() { done <- c.Wait() }()
			select {
			case <-done:
			case <-time.After(time.Second):
				t.Fatal("timeout not enforced")
			}
		})
	}
}
func TestDisconnectReleasesCapacity(t *testing.T) {
	addr, key := startTestServer(t, func(c *config) { c.MaxConnections = 1 })
	c := client(t, addr, key)
	_ = c.Close()
	time.Sleep(40 * time.Millisecond)
	c2 := client(t, addr, key)
	_ = c2.Close()
}
func TestOutputLimit(t *testing.T) {
	a, b := net.Pipe()
	defer b.Close()
	released := false
	c := &guardedConn{Conn: a, release: func() { released = true }}
	_, e := c.Write(make([]byte, 2*1024*1024+1))
	if e == nil || !released {
		t.Fatal("output limit failed")
	}
}

func TestBannerResponsiveMenu(t *testing.T) {
	c, _ := loadContent()
	for _, width := range []int{24, 40, 49, 60, 80, 160} {
		for _, height := range []int{10, 16, 22, 23, 24, 40} {
			m := newModel(c, lipgloss.NewStyle())
			m.width, m.height = width, height
			for cursor := 0; cursor < 8; cursor++ {
				m.cursor = cursor
				view := ansi.Strip(m.View())
				lines := strings.Split(strings.TrimSuffix(view, "\n"), "\n")
				if len(lines) > height {
					t.Fatalf("height overflow at %dx%d: %d", width, height, len(lines))
				}
				for _, line := range lines {
					if ansi.StringWidth(line) > width {
						t.Fatalf("width overflow at %dx%d", width, height)
					}
				}
				if !strings.Contains(view, fmt.Sprintf("> %d", cursor+1)) {
					t.Fatalf("selection missing at %dx%d", width, height)
				}
				hasBanner := strings.Contains(view, strings.Split(nameBanner, "\n")[0])
				if hasBanner != (width >= 49 && height >= 23) {
					t.Fatalf("incorrect banner at %dx%d", width, height)
				}
			}
			m.page = 0
			if strings.Contains(m.View(), nameBanner) {
				t.Fatal("banner takes space from section content")
			}
		}
	}
	if strings.Contains(plainContent(c), "\x1b") || !strings.Contains(plainContent(c), nameBanner) {
		t.Fatal("plain banner missing or styled")
	}
}
