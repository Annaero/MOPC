package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventType string

const (
	Offline string = "offline"
	Online         = "online"
	Duel           = "duel"
)

type NewEvent struct {
	Name        string    `json:"name" bson:"name,omitempty" validate:"required"`
	Type        EventType `json:"type" bson:"type,omitempty" validate:"required"`
	Description string    `json:"description" bson:"description"`
	StartDate   Date      `json:"startDate" bson:"startDate,omitempty" validate:"required,ltecsfield=EndDate"`
	EndDate     Date      `json:"endDate" bson:"endDate"`
}

type EventID struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

type Event struct {
	EventID
	NewEvent
}
