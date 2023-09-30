package events

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/Annaero/MOPC/goberry/mocks/database"
)

var (
	mockDB       = database.MockDB{}
	eventJSON    = `{}`
	newEventJSON = `{"name":"Test inserted event", "type":"offine", "startDate":"2022-12-12"}`
	eventIdJSON  = `{"id":"650ffd7eb0756fac7f2de46a"}`
)

func TestCreateEvent(t *testing.T) {
	e := echo.New()
	h := &EventsHandler{db: mockDB, validator: validator.New()}
	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPost, groupURL, strings.NewReader(newEventJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.postEvent(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, eventIdJSON, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestGetEvents(t *testing.T) {
	e := echo.New()
	h := &EventsHandler{db: mockDB, validator: validator.New()}

	testCases := map[string]struct {
		params     map[string]string
		statusCode int
		response   interface{}
	}{
		"Good request": {
			map[string]string{
				"start_date": "2021-01-01", "end_date": "2021-01-02",
			},
			http.StatusOK,
			nil,
		},
		"Empty params": {
			map[string]string{},
			http.StatusBadRequest,
			nil,
		},
	}

	for tc, tp := range testCases {
		req, _ := http.NewRequest("GET", groupURL, nil)
		q := req.URL.Query()
		for k, v := range tp.params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if assert.NoError(t, h.getEvents(c)) {
			assert.Equal(t, tp.statusCode, rec.Code, tc, rec.Body)
			if tp.response != nil {
				assert.Equal(t, tp.response, strings.TrimSuffix(rec.Body.String(), "\n"))
			}
		}
	}
}
