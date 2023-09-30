package database

import (
	"time"

	"github.com/Annaero/MOPC/goberry/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockDB struct {
	events []models.Event
}

func (mdb MockDB) GetEvents(startDate time.Time, endDate time.Time) ([]models.Event, error) {
	return mdb.events, nil
}

func (mdb MockDB) GetEvent(eventId primitive.ObjectID) (*models.Event, error) {
	event := mdb.events[0]
	return &event, nil
}

func (mdb MockDB) CreateEvent(event models.Event) (*primitive.ObjectID, error) {
	oid, _ := primitive.ObjectIDFromHex("650ffd7eb0756fac7f2de46a")
	return &oid, nil
}
