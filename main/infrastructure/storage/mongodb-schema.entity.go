package storage

import (
	"fakerAPI/main/domain/properties"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoSchema struct {
	Id         primitive.ObjectID    `bson:"_id"`
	Name       string                `bson:"name"`
	UserId     string                `bson:"user_id"`
	Properties []MongoSchemaProperty `bson:"properties"`
}
type MongoSchemaProperty struct {
	Name       string                `bson:"name"`
	Type       properties.Type       `bson:"type"`
	Min        int                   `bson:"min,omitempty"`
	Max        int                   `bson:"max,omitempty"`
	Rate       float32               `bson:"rate,omitempty"`
	From       time.Time             `bson:"from,omitempty"`
	To         time.Time             `bson:"to,omitempty"`
	Element    *MongoSchemaProperty  `bson:"element,omitempty"`
	RangeSize  [2]int                `bson:"rangeSize,omitempty"`
	Properties []MongoSchemaProperty `bson:"properties,omitempty"`
	Enums      []string              `bson:"enums,omitempty"`
}
