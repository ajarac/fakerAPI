package storage

import (
	"fakerAPI/main/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoSchema struct {
	Id         primitive.ObjectID     `bson:"_id"`
	Name       string                 `bson:"name"`
	Properties []*MongoSchemaProperty `bson:"properties"`
}
type MongoSchemaProperty struct {
	Name       string                 `bson:"name"`
	Type       domain.Type            `bson:"type"`
	Min        int                    `bson:"min,omitempty"`
	Max        int                    `bson:"max,omitempty"`
	Rate       float64                `bson:"rate,omitempty"`
	Element    *MongoSchemaProperty   `bson:"element,omitempty"`
	RangeSize  []int                  `bson:"rangeSize,omitempty"`
	Properties []*MongoSchemaProperty `bson:"properties,omitempty"`
}
