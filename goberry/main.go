package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type EventType string

const (
	Offline string = "offline"
	Online         = "online"
	Duel           = "duel"
)

type Event struct {
	ID          uuid.UUID `json:"id" validate:"required,uuid"`
	Name        string    `json:"name" validate:"required"`
	Type        EventType `json:"type" validate:"required"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate" validate:"required"`
	EndDate     time.Time `json:"endDate"`
}

type IError struct {
	Field string
	Tag   string
	Value string
}

// var Validator = validator.New()

// func ValidateAddDog(c *fiber.Ctx) error {
// 	var errors []*IError
// 	body := new(Dog)
// 	c.BodyParser(&body)

// 	err := Validator.Struct(body)
// 	if err != nil {
// 		for _, err := range err.(validator.ValidationErrors) {
// 			var el IError
// 			el.Field = err.Field()
// 			el.Tag = err.Tag()
// 			el.Value = err.Param()
// 			errors = append(errors, &el)
// 		}
// 		return c.Status(fiber.StatusBadRequest).JSON(errors)
// 	}
// 	return c.Next()
// }

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/events", func(w http.ResponseWriter, r *http.Request) {
		var events = []Event{
			{Name: "Some Name",
				Type:        Online,
				ID:          uuid.New(),
				Description: "",
				StartDate:   time.Now(),
				EndDate:     time.Now()}}

		render.JSON(w, r, events)
	})
	http.ListenAndServe(":3000", r)
}
