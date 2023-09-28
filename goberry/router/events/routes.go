package events

import "github.com/go-chi/chi/v5"

const (
	groupURL = "/events"
)

// Routes function to create router.
func Routes(m *chi.Mux) {
	// Create group.
	m.Route(groupURL, func(r chi.Router) {
		r.Get("/", getEvents) // get status route
	})
}
