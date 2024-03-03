package http

import (
	"log"
	"pdfy/internal/infrastructure/http/handlers"
	"pdfy/internal/infrastructure/http/routes"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
}

func NewServer(allHandlers *handlers.HandlerContainer) (*Server, error) {
	log.Println("starting new http server")
	router, err := routes.RegisterRoutes(allHandlers)
	if err != nil {
		return nil, err
	}

	return &Server{
		Router: router,
	}, nil
}
