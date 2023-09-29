package events

import (
	"github.com/Annaero/MOPC/goberry/models/database"
	"github.com/labstack/echo/v4"
)

const (
	groupURL = "/api/events"
)

// Routes function to create router.
func Routes(e *echo.Echo, db database.EventsDB) {
	// Create group.
	g := e.Group(groupURL)
	handler := EventsHandler{db: db}

	g.GET("", handler.getEvents)
	g.GET("/:event_id", handler.getEventByID)
}
