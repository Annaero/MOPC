package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
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
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Thank god it works 🙏")
	})

	app.Get("/api/events", func(c *fiber.Ctx) error {

		var events = []Event{
			Event{Name: "Some Name",
				Type:        Online,
				ID:          uuid.New(),
				Description: "",
				StartDate:   time.Now(),
				EndDate:     time.Now()}}

		return c.JSON(fiber.Map{
			"error":  false,
			"msg":    nil,
			"count":  len(events),
			"events": events,
		})
	})

	app.Get("/api/events/{event_id}", func(c *fiber.Ctx) error {

		var events = []Event{
			Event{Name: "Some Name",
				Type:        Online,
				ID:          uuid.New(),
				Description: "",
				StartDate:   time.Now(),
				EndDate:     time.Now()}}

		return c.JSON(fiber.Map{
			"error":  false,
			"msg":    nil,
			"count":  len(events),
			"events": events,
		})
	})

	app.Listen(":8000")
}
