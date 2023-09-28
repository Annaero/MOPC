package events

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	models "github.com/Annaero/MOPC/goberry/models"
)

func getEventByID(c echo.Context) error {
	eventIdParam := c.Param("event_id")
	eventId, err := primitive.ObjectIDFromHex(eventIdParam)
	if err != nil {
		echo.NewHTTPError(echo.ErrNotFound.Code)
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	collection := client.Database("events").Collection("mopc_event")
	filter := bson.D{{"_id", eventId}}
	var result models.Event
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(echo.ErrNotFound.Code)
		}
		panic(err)
	}

	return c.JSON(http.StatusOK, result)

}

func getEvents(c echo.Context) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	var startDate models.SimpleDate
	var endDate models.SimpleDate

	// creates query params binder that stops binding at first error
	binderErrors := echo.QueryParamsBinder(c).
		MustBindUnmarshaler("start_date", &startDate).
		MustBindUnmarshaler("end_date", &endDate).
		BindErrors() // returns all binding error

	if binderErrors != nil {
		errorMessage := ""
		for _, err := range binderErrors {
			errorMessage += err.Error() + " "
		}
		return c.String(http.StatusBadRequest, errorMessage)
	}

	collection := client.Database("events").Collection("mopc_event")
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"startDate", bson.D{{"$lt", time.Time(endDate)}}}},
				bson.D{{"endDate", bson.D{{"$gt", time.Time(startDate)}}}},
			},
		},
	}
	results := &[]models.Event{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.TODO(), results); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, results)
}
