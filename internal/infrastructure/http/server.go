package http

import (
	"pdfy/internal/infrastructure/http/handlers"
	"pdfy/internal/infrastructure/http/routes"
	"pdfy/pkg/logger"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
}

func NewServer(allHandlers *handlers.HandlerContainer) (*Server, error) {
	logger.Info("starting the server", "port", "8080")
	router, err := routes.RegisterRoutes(allHandlers)
	if err != nil {
		return nil, err
	}

	return &Server{
		Router: router,
	}, nil
}
