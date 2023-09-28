package router

import (
	"net/http"

	"github.com/Annaero/MOPC/goberry/router/events"
	"github.com/go-chi/chi/v5"
)

// GetRoutes function for getting routes.
func GetRoutes(m *chi.Mux) {
	events.Routes(m)          // health check routes
	m.NotFound(http.NotFound) // not found routes
}
