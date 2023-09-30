package events

import (
	"github.com/Annaero/MOPC/goberry/models/database"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

const (
	groupURL = "/api/events"
)

func Routes(e *echo.Echo, db database.EventsDB) {
	g := e.Group(groupURL)

	handler := EventsHandler{db: db, validator: validator.New()}

	g.GET("", handler.getEvents)
	g.GET("/:event_id", handler.getEventByID)
	g.POST("", handler.postEvent)
}
