package router

import (
	"github.com/Annaero/MOPC/goberry/router/events"
	"github.com/labstack/echo/v4"
)

// GetRoutes function for getting routes.
func GetRoutes(e *echo.Echo) {
	events.Routes(e) // health check routes
	e.RouteNotFound("/*", echo.NotFoundHandler)
}
