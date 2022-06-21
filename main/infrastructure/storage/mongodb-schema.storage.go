package storage

import (
	"context"
	"fakerAPI/main/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBSchemaStorage struct {
	client *mongo.Collection
}

func (m *MongoDBSchemaStorage) GetNextId() string {
	return primitive.NewObjectID().Hex()
}

func (m *MongoDBSchemaStorage) Create(schema *domain.Schema) (*domain.Schema, error) {
	mongoSchema, err := fromDomain(schema)
	if err != nil {
		return nil, err
	}
	_, err = m.client.InsertOne(context.Background(), mongoSchema)
	if err != nil {
		return nil, err
	}
	return schema, nil
}

func (m *MongoDBSchemaStorage) GetById(id string) (*domain.Schema, bool, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, false, err
	}
	var mongoSchema MongoSchema
	cursor := m.client.FindOne(context.Background(), bson.D{{"_id", objectId}})
	err = cursor.Decode(&mongoSchema)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	return toDomain(&mongoSchema), true, nil
}

func (m *MongoDBSchemaStorage) GetAll() ([]*domain.Schema, error) {
	ctx := context.Background()
	cursor, err := m.client.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var schemas []*domain.Schema
	for cursor.Next(ctx) {
		var mongoSchema MongoSchema
		err = cursor.Decode(&mongoSchema)
		if err != nil {
			return nil, err
		}
		schemas = append(schemas, toDomain(&mongoSchema))
	}
	return schemas, nil
}

func (m *MongoDBSchemaStorage) Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = m.client.DeleteOne(context.Background(), bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}

func NewMongoDBSchemaStorage(database *mongo.Database) *MongoDBSchemaStorage {
	return &MongoDBSchemaStorage{client: database.Collection("schemas")}
}
