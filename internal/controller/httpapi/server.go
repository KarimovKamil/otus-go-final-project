package httpapi

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KarimovKamil/otus-go-final-project/internal/config"
)

type Server struct {
	server  *http.Server
	handler http.Handler
	config  *config.Config
}

func NewServer(handler http.Handler, config *config.Config) *Server {
	return &Server{
		config:  config,
		handler: handler,
	}
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:         s.config.Server.Port,
		Handler:      s.handler,
		ReadTimeout:  time.Duration(s.config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.config.Server.WriteTimeout) * time.Second,
	}
	return s.server.ListenAndServe()
}

func (s *Server) ShutdownService(c chan os.Signal) {
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.server.Shutdown(ctx)
}
