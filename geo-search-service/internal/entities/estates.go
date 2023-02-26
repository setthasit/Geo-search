package entities

import (
	"time"
)

const (
	ESTATES_COLLECTION = "estates"
	ESTATES_INDEX      = "estates"
)

type Estate struct {
	Title    string   `json:"title" bson:"title"`
	Location Location `json:"location" bson:"location"`

	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
}

type Location struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lon float64 `json:"lon" bson:"lon"`
}
