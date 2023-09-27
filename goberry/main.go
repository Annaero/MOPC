package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ID          primitive.ObjectID `bson:"_id"json:"id" validate:"required,uuid"`
	Name        string             `json:"name" validate:"required"`
	Type        EventType          `json:"type" validate:"required"`
	Description string             `json:"description"`
	StartDate   time.Time          `json:"startDate" validate:"required"`
	EndDate     time.Time          `json:"endDate"`
}

type RequestError struct {
	Code        int    `json:"error"`
	Description string `json:"description"`
}

func (e *RequestError) Render(w http.ResponseWriter, r *http.Request) {
	render.Status(r, e.Code)
	render.JSON(w, r, e)
}

type IError struct {
	Field string
	Tag   string
	Value string
}

const eventDateLayout = "2006-01-02"

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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/events", func(w http.ResponseWriter, r *http.Request) {

		startDateParam := r.URL.Query().Get("start_date")
		endDateParam := r.URL.Query().Get("end_date")
		if startDateParam == "" || endDateParam == "" {
			var my_error = RequestError{
				Code:        400,
				Description: "startDate is missing",
			}
			my_error.Render(w, r)
			return
		}

		startDate, startDateErr := time.Parse(eventDateLayout, startDateParam)
		endDate, endDateErr := time.Parse(eventDateLayout, endDateParam)
		if startDateErr != nil || endDateErr != nil {
			var errorDescription = ""
			if startDateErr != nil {
				errorDescription += startDateErr.Error()
			}
			if endDateErr != nil {
				errorDescription += endDateErr.Error()
			}
			var my_error = RequestError{
				Code:        422,
				Description: errorDescription,
			}
			my_error.Render(w, r)
			return
		}

		filter := bson.D{
			{"$and",
				bson.A{
					bson.D{{"startDate", bson.D{{"$lt", endDate}}}},
					bson.D{{"endDate", bson.D{{"$gt", startDate}}}},
				},
			},
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
