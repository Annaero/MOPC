package router

import (
	"github.com/Annaero/MOPC/goberry/models/database"
	"github.com/Annaero/MOPC/goberry/router/events"
	"github.com/labstack/echo/v4"
)

// GetRoutes function for getting routes.
func GetRoutes(e *echo.Echo, db database.GoberryDB) {
	events.Routes(e, db) // health check routes
	e.RouteNotFound("/*", echo.NotFoundHandler)
}
