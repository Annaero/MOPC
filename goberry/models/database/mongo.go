package database

import (
	"context"
	"errors"
	"time"

	"github.com/Annaero/MOPC/goberry/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	Client *mongo.Client
}

func (mdb MongoDB) GetEvents(startDate time.Time, endDate time.Time) ([]models.Event, error) {
	collection := mdb.Client.Database("events").Collection("mopc_event")

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"startDate", bson.D{{"$lt", endDate}}}},
				bson.D{{"endDate", bson.D{{"$gt", startDate}}}},
			},
		},
	}

	results := []models.Event{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, err
}

func (mdb MongoDB) GetEvent(eventId primitive.ObjectID) (*models.Event, error) {
	collection := mdb.Client.Database("events").Collection("mopc_event")

	filter := bson.D{{Key: "_id", Value: eventId}}

	var result models.Event
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (mdb MongoDB) CreateEvent(event models.Event) (*primitive.ObjectID, error) {
	collection := mdb.Client.Database("events").Collection("mopc_event")

	result, err := collection.InsertOne(context.TODO(), &event)
	if err != nil {
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return &oid, nil
	}

	return nil, errors.New("OID is not OK")
}
