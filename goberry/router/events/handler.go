package events

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Annaero/MOPC/goberry/models"
	"github.com/Annaero/MOPC/goberry/models/database"
)

type EventsHandler struct {
	db        database.EventsDB
	validator *validator.Validate
}

func (eh EventsHandler) postEvent(c echo.Context) error {
	var event models.Event
	if err := c.Bind(&event); err != nil {
		log.Error("Event bind error:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := eh.validator.Struct(event); err != nil {
		log.Error("Event validation error:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Debug(event)

	oid, err := eh.db.CreateEvent(event)
	if err != nil {
		panic(err)
	}

	response := models.EventID{ID: *oid}
	return c.JSON(http.StatusCreated, response)
}

func (eh EventsHandler) getEventByID(c echo.Context) error {
	eventIdParam := c.Param("event_id")
	eventId, err := primitive.ObjectIDFromHex(eventIdParam)
	if err != nil {
		echo.NewHTTPError(echo.ErrNotFound.Code)
	}

	event, err := eh.db.GetEvent(eventId)
	if err != nil {
		panic(err)
	}

	if event != nil {
		return c.JSON(http.StatusOK, event)
	}

	return echo.NewHTTPError(echo.ErrNotFound.Code)
}

func (eh EventsHandler) getEvents(c echo.Context) error {
	var startDate time.Time
	var endDate time.Time

	binderErrors := echo.QueryParamsBinder(c).
		MustTime("start_date", &startDate, time.DateOnly).
		MustTime("end_date", &endDate, time.DateOnly).
		BindErrors() // returns all binding error

	if binderErrors != nil {
		errorMessage := ""
		for _, err := range binderErrors {
			errorMessage += err.Error() + " "
		}
		return c.String(http.StatusBadRequest, errorMessage)
	}

	if endDate.Before(startDate) {
		s := fmt.Sprintf("end_date: %s is before start_date: %s", endDate, startDate)
		return c.String(http.StatusBadRequest, s)
	}

	events, err := eh.db.GetEvents(startDate, endDate)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, events)
}
