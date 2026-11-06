package main

import (
	"context"
	"errors"
	"github.com/charmbracelet/ssh"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if os.Geteuid() == 0 {
		log.Fatal("Run as an unprivileged user, not root")
	}
	c := defaultConfig()
	if a := os.Getenv("PORTFOLIO_ADDR"); a != "" {
		c.Address = a
	}
	if k := os.Getenv("PORTFOLIO_HOST_KEY"); k != "" {
		c.HostKey = k
	}
	content, err := loadContent()
	if err != nil {
		log.Fatal(err)
	}
	s, err := makeServer(c, content)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.Listen("tcp", c.Address)
	if err != nil {
		log.Fatal(err)
	}
	limited := &limitedListener{listener, newLimiter(c.MaxConnections), c.Handshake, c.Idle, c.Duration}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	done := make(chan error, 1)
	go func() { done <- s.Serve(limited) }()
	log.Printf("Portfolio listening on %s", listener.Addr())
	select {
	case err = <-done:
		if err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Fatal("SSH server stopped")
		}
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if s.Shutdown(timeout) != nil {
			_ = s.Close()
		}
	}
}
