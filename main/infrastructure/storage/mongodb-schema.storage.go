package storage

import (
	"context"
	"fakerAPI/main/config"
	"fakerAPI/main/domain"
	"fmt"
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

func (m *MongoDBSchemaStorage) Create(context context.Context, schema *domain.Schema) (*domain.Schema, error) {
	user := getUserByContext(context)
	mongoSchema, err := fromDomain(schema, user)
	if err != nil {
		return nil, err
	}
	_, err = m.client.InsertOne(context, mongoSchema)
	if err != nil {
		return nil, err
	}
	return schema, nil
}

func (m *MongoDBSchemaStorage) GetById(context context.Context, id string) (*domain.Schema, bool, error) {
	user := getUserByContext(context)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, false, err
	}
	var mongoSchema MongoSchema
	cursor := m.client.FindOne(context, bson.D{{Key: "$and", Value: bson.A{bson.D{{"_id", objectId}}, bson.D{{"user", user}}}}})
	err = cursor.Decode(&mongoSchema)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	return toDomain(&mongoSchema), true, nil
}

func (m *MongoDBSchemaStorage) GetAll(context context.Context) ([]*domain.Schema, error) {
	user := getUserByContext(context)
	cursor, err := m.client.Find(context, bson.D{{"user", user}})
	if err != nil {
		return nil, err
	}
	var schemas []*domain.Schema
	for cursor.Next(context) {
		var mongoSchema MongoSchema
		err = cursor.Decode(&mongoSchema)
		if err != nil {
			return nil, err
		}
		schemas = append(schemas, toDomain(&mongoSchema))
	}
	return schemas, nil
}

func (m *MongoDBSchemaStorage) Delete(context context.Context, id string) error {
	user := getUserByContext(context)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = m.client.DeleteOne(context, bson.D{{Key: "$and", Value: bson.A{bson.D{{"_id", objectId}}, bson.D{{"user", user}}}}})
	if err != nil {
		return err
	}
	return nil
}

func getUserByContext(context context.Context) string {
	return fmt.Sprintf("%s", context.Value(config.UserContext))
}

func NewMongoDBSchemaStorage(database *mongo.Database) *MongoDBSchemaStorage {
	mod := mongo.IndexModel{
		Keys:    bson.M{"user": 1},
		Options: nil,
	}
	collection := database.Collection("schemas")
	collection.Indexes().CreateOne(context.Background(), mod)
	return &MongoDBSchemaStorage{client: collection}
}
