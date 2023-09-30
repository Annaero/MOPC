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
	Name        string    `json:"name" validate:"required"`
	Type        EventType `json:"type" validate:"required"`
	Description string    `json:"description"`
	StartDate   Date      `json:"startDate" validate:"required"`
	EndDate     Date      `json:"endDate"`
}

type EventID struct {
	ID primitive.ObjectID `bson:"_id" json:"id" validate:"required,uuid"`
}

type Event struct {
	EventID
	NewEvent
}
