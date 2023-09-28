package events

import (
	"github.com/labstack/echo/v4"
)

const (
	groupURL = "/api/events"
)

// Routes function to create router.
func Routes(e *echo.Echo) {
	// Create group.
	g := e.Group(groupURL)
	g.GET("", getEvents)
}
