package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Annaero/MOPC/goberry/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client           *mongo.Client
	eventsCollection *mongo.Collection
}

func ConnectMongDB(connectionString string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	log.Print("Connected to MongoDB")

	opts := options.Collection().SetRegistry(models.RegisterDate())
	eventsCollection := client.Database("events").Collection("mopc_event", opts)

	return &MongoDB{client: client, eventsCollection: eventsCollection}, nil
}

func (mdb MongoDB) GetEvents(startDate time.Time, endDate time.Time) ([]models.Event, error) {
	filter :=
		bson.D{{"$or",
			bson.A{
				bson.D{{"$and",
					bson.A{
						bson.D{{"endDate", bson.D{{"$exists", true}}}},
						bson.D{{"startDate", bson.D{{"$lte", endDate}}}},
						bson.D{{"endDate", bson.D{{"$gte", startDate}}}},
					},
				}},
				bson.D{{"$and",
					bson.A{
						bson.D{{"endDate", bson.D{{"$exists", false}}}},
						bson.D{{"startDate", bson.D{{"$lte", endDate}}}},
						bson.D{{"startDate", bson.D{{"$gte", startDate}}}},
					},
				}},
			},
		}}

	results := []models.Event{}
	cursor, err := mdb.eventsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, err
}

func (mdb MongoDB) GetEvent(eventId primitive.ObjectID) (*models.Event, error) {
	filter := bson.D{{Key: "_id", Value: eventId}}

	var result models.Event
	err := mdb.eventsCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (mdb MongoDB) CreateEvent(event models.Event) (*primitive.ObjectID, error) {
	result, err := mdb.eventsCollection.InsertOne(context.TODO(), &event)
	if err != nil {
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return &oid, nil
	}

	return nil, errors.New("OID is not OK")
}
