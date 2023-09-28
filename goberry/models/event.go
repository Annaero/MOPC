package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventType string

const (
	Offline string = "offline"
	Online         = "online"
	Duel           = "duel"
)

type Event struct {
	ID          primitive.ObjectID `bson:"_id" json:"id" validate:"required,uuid"`
	Name        string             `json:"name" validate:"required"`
	Type        EventType          `json:"type" validate:"required"`
	Description string             `json:"description"`
	StartDate   time.Time          `json:"startDate" validate:"required"`
	EndDate     time.Time          `json:"endDate"`
}
