package storage

import (
	"context"
	"fakerAPI/main/config"
	"fakerAPI/main/domain"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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
	_, err = m.client.InsertOne(context, mongoSchema)
	if err != nil {
		return nil, err
	}
	return schema, nil
}
func (m *MongoDBSchemaStorage) GetByName(context context.Context, name string) (*domain.Schema, bool, error) {
	user := getUserByContext(context)
	filter := bson.D{{Key: "$and", Value: bson.A{bson.D{{"name", name}}, bson.D{{"user_id", user}}}}}
	return m.getOneSchemaByFilter(context, filter)
}

func (m *MongoDBSchemaStorage) GetById(context context.Context, id string) (*domain.Schema, bool, error) {
	user := getUserByContext(context)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, false, err
	}
	filter := bson.D{{Key: "$and", Value: bson.A{bson.D{{"_id", objectId}}, bson.D{{"user_id", user}}}}}
	return m.getOneSchemaByFilter(context, filter)
}

func (m *MongoDBSchemaStorage) GetAll(context context.Context) ([]*domain.Schema, error) {
	user := getUserByContext(context)
	cursor, err := m.client.Find(context, bson.D{{"user_id", user}})
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
		schema, err := toDomain(&mongoSchema)
		if err != nil {
			return nil, err
		}
		schemas = append(schemas, schema)
	}
	return schemas, nil
}

func (m *MongoDBSchemaStorage) Delete(context context.Context, id string) error {
	user := getUserByContext(context)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = m.client.DeleteOne(context, bson.D{{Key: "$and", Value: bson.A{bson.D{{"_id", objectId}}, bson.D{{"user_id", user}}}}})
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDBSchemaStorage) getOneSchemaByFilter(context context.Context, filter interface{}) (*domain.Schema, bool, error) {
	var mongoSchema MongoSchema
	cursor := m.client.FindOne(context, filter)
	err := cursor.Decode(&mongoSchema)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	schema, err := toDomain(&mongoSchema)
	if err != nil {
		return nil, false, err
	}
	return schema, true, nil
}

func getUserByContext(context context.Context) string {
	return fmt.Sprintf("%s", context.Value(config.UserContext))
}

func NewMongoDBSchemaStorage(database *mongo.Database) *MongoDBSchemaStorage {
	collection := database.Collection("schemas")
	createIndexes(collection)
	return &MongoDBSchemaStorage{client: collection}
}

func createIndexes(collection *mongo.Collection) {
	models := []mongo.IndexModel{{
		Keys: bson.M{"user_id": 1},
	}, {
		Keys: bson.M{"name": 1},
	}, {
		Keys: bson.D{{Key: "user_id", Value: 1}, {Key: "name", Value: 1}},
	}}
	_, err := collection.Indexes().CreateMany(context.Background(), models)
	if err != nil {
		log.Panicf("Error creating indexes %s\n", err.Error())
	}
}
