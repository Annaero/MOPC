package events

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Annaero/MOPC/goberry/models/database"
)

const SIMPLE_DATE_LAYOUT = "2006-01-02"

type EventsHandler struct {
	db database.EventsDB
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
		Time("start_date", &startDate, SIMPLE_DATE_LAYOUT).
		Time("end_date", &endDate, SIMPLE_DATE_LAYOUT).
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
