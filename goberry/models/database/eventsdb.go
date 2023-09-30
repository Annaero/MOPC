package database

import (
	"time"

	"github.com/Annaero/MOPC/goberry/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventsDB interface {
	GetEvents(startDate time.Time, endDate time.Time) ([]models.Event, error)
	GetEvent(eventId primitive.ObjectID) (*models.Event, error)
	CreateEvent(event models.Event) (*primitive.ObjectID, error)
}

type GoberryDB interface {
	EventsDB
}
