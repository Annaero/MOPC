package events

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	helpers "github.com/Annaero/MOPC/goberry/helpers/request"
	models "github.com/Annaero/MOPC/goberry/models"
)

const eventDateLayout = "2006-01-02"

func getEvents(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	eventsParamsNames := map[string]bool{"start_date": true, "end_date": true}
	eventsParams, err := helpers.GetRequestParams(r, eventsParamsNames)
	if err != nil {
		var startDateMissingErr = helpers.RequestError{
			Code:        400,
			Description: err.Error(),
		}
		startDateMissingErr.Render(w, r)
		return
	}

	startDate, startDateErr := time.Parse(eventDateLayout, eventsParams["start_date"])
	endDate, endDateErr := time.Parse(eventDateLayout, eventsParams["end_date"])
	if startDateErr != nil || endDateErr != nil {
		var errorDescription = ""
		if startDateErr != nil {
			errorDescription += startDateErr.Error()
			errorDescription += " "
		}
		if endDateErr != nil {
			errorDescription += endDateErr.Error()
		}
		var my_error = helpers.RequestError{
			Code:        422,
			Description: errorDescription,
		}
		my_error.Render(w, r)
		return
	}

	collection := client.Database("events").Collection("mopc_event")
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"startDate", bson.D{{"$lt", endDate}}}},
				bson.D{{"endDate", bson.D{{"$gt", startDate}}}},
			},
		},
	}
	var results []models.Event
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	render.JSON(w, r, results)
	return
}
