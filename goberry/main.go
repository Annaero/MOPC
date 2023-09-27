package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

type UnprocessableEntity struct {
	Error int    `json:"error"`
	Text  string `json:"error_text"`
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("events").Collection("mopc_event")
	filter := bson.D{{}}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/events", func(w http.ResponseWriter, r *http.Request) {

		startDateParam := r.URL.Query().Get("startDate")
		if startDateParam == "" {
			var my_error = UnprocessableEntity{
				Error: 400,
				Text:  "startDateParam",
			}
			render.Status(r, http.StatusUnprocessableEntity)
			render.JSON(w, r, my_error)
			return
		}

		_, err := time.Parse("2006-01-02", startDateParam)
		if err != nil {
			// app.logger.Info().Msgf("can not parse ID: %v", id)
			var my_error = UnprocessableEntity{
				Error: 422,
				Text:  err.Error(),
			}
			render.Status(r, http.StatusUnprocessableEntity)
			render.JSON(w, r, my_error)
			return
		}

		var results []Event
		cursor, err := collection.Find(context.TODO(), filter)
		if err != nil {
			panic(err)
		}
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}
		// for _, result := range results {
		// 	res, _ := json.Marshal(result)
		// 	fmt.Println(string(res))
		// }

		// var events = []Event{
		// 	{Name: "Some Name",
		// 		Type:        Online,
		// 		ID:          uuid.New(),
		// 		Description: "",
		// 		StartDate:   time.Now(),
		// 		EndDate:     time.Now()}}

		render.JSON(w, r, results)
	})
	http.ListenAndServe(":3000", r)
}
