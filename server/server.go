package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"net"
	"net/http"
	"strconv"
	"time"
)

type MathSrv struct {
	Address string
	Mux chi.Router
	Server *http.Server
}

func NewServer(host string, port int) *MathSrv {
	address := net.JoinHostPort(host, strconv.Itoa(port))
	mux := chi.NewMux()

	return &MathSrv {
		Address: address,
		Mux: mux,
		Server: &http.Server{
			Addr: address,
			Handler: mux,
			ReadTimeout:       5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      5 * time.Second,
			IdleTimeout:       5 * time.Second,
		},
	}
}

func (s *MathSrv) setupRoutes() {
	// no-op
}

func (s *MathSrv) Start() error {
	s.setupRoutes()

	fmt.Println("Starting on", s.Address)
	if err := s.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("error starting server: %w", err)
	}
	return nil
}

func (s *MathSrv) Stop() error {
	fmt.Println("Stopping server ", s.Address)
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		return fmt.Errorf("error shutting down server: %w", err)
	}

	return nil
}