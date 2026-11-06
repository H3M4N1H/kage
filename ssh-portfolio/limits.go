package main

import (
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type bucket struct {
	start            time.Time
	attempts, active int
}
type limiter struct {
	mu          sync.Mutex
	entries     map[string]*bucket
	active, max int
}

func newLimiter(max int) *limiter { return &limiter{entries: make(map[string]*bucket), max: max} }
func (l *limiter) acquire(ip string, now time.Time) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	for k, b := range l.entries {
		if b.active == 0 && now.Sub(b.start) >= time.Minute {
			delete(l.entries, k)
		}
	}
	b := l.entries[ip]
	if b == nil {
		if len(l.entries) >= 4096 {
			return false
		}
		b = &bucket{start: now}
		l.entries[ip] = b
	}
	if now.Sub(b.start) >= time.Minute {
		b.start = now
		b.attempts = 0
	}
	if b.attempts >= 6 {
		return false
	}
	b.attempts++
	if b.active >= 2 || l.active >= l.max {
		return false
	}
	b.active++
	l.active++
	return true
}
func (l *limiter) release(ip string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if b := l.entries[ip]; b != nil && b.active > 0 {
		b.active--
		l.active--
	}
}

type guardedConn struct {
	net.Conn
	once          sync.Once
	release       func()
	input, output atomic.Int64
	opened        atomic.Bool
	start         time.Time
	handshake     time.Duration
	idle          time.Duration
}

func (c *guardedConn) Close() error { err := c.Conn.Close(); c.once.Do(c.release); return err }

// Deadlines are owned by this wrapper. The SSH library calls SetDeadline
// during writes; forwarding that call would erase an in-flight read deadline.
func (c *guardedConn) SetDeadline(time.Time) error { return nil }

func (c *guardedConn) Read(p []byte) (int, error) {
	deadline := time.Now().Add(c.idle)
	if !c.opened.Load() {
		deadline = minTime(deadline, c.start.Add(c.handshake))
	}
	_ = c.Conn.SetReadDeadline(deadline)
	n, e := c.Conn.Read(p)
	if c.input.Add(int64(n)) > 256*1024 {
		_ = c.Close()
		return 0, errors.New("input budget exceeded")
	}
	return n, e
}
func (c *guardedConn) Write(p []byte) (int, error) {
	if c.output.Add(int64(len(p))) > 2*1024*1024 {
		_ = c.Close()
		return 0, errors.New("output budget exceeded")
	}
	_ = c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	return c.Conn.Write(p)
}
func minTime(a, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}

type limitedListener struct {
	net.Listener
	limits                    *limiter
	handshake, idle, duration time.Duration
}

func (l *limitedListener) Accept() (net.Conn, error) {
	for {
		c, e := l.Listener.Accept()
		if e != nil {
			return nil, e
		}
		ip, _, e := net.SplitHostPort(c.RemoteAddr().String())
		if e != nil || !l.limits.acquire(ip, time.Now()) {
			_ = c.Close()
			continue
		}
		g := &guardedConn{Conn: c, release: func() { l.limits.release(ip) }, start: time.Now(), handshake: l.handshake, idle: l.idle}
		timer := time.AfterFunc(l.duration, func() { _ = g.Close() })
		release := g.release
		g.release = func() { timer.Stop(); release() }
		return g, nil
	}
}
