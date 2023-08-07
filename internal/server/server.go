package server

import (
	"context"
	"fmt"

	"github.com/lakhansamani/create-go-graphql-server/internal/router"
	"github.com/rs/zerolog"
)

// Config is the configuration for a server
type Config struct {
	// Port is the port to listen on
	Port int
}

// Server is a GraphQL server
type Server struct {
	Config
	log zerolog.Logger
}

// New creates a new server
func New(log zerolog.Logger, cfg Config) (*Server, error) {
	return &Server{
		Config: cfg,
		log:    log,
	}, nil
}

// Run listens and servers the GraphQL server
func (s *Server) Run(ctx context.Context) error {
	router := router.New()
	s.log.Debug().Int("port", s.Port).Msg("Listening")
	router.Run(fmt.Sprintf(":%d", s.Port))
	return nil
}
