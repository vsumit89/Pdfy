package routes

import (
	"pdfy/internal/infrastructure/http/handlers"

	"github.com/go-chi/chi/v5"
)

// RegisterRoutes binds handlers to there corresponding routes
func RegisterRoutes(allHandlers *handlers.HandlerContainer) (*chi.Mux, error) {
	mux := chi.NewRouter()

	apiV1Router := mountAPIv1Router(allHandlers)

	// adding prefix /api/v1 to the routes
	mux.Mount("/api/v1", apiV1Router)
	return mux, nil
}

func mountAPIv1Router(handlers *handlers.HandlerContainer) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", handlers.TemplateHandler.CreateTemplate)
	return router
}
