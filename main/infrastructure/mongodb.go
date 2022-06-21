package infrastructure

import (
	"context"
	"fakerAPI/main/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const DatabaseName = "faker"

func GetDatabase(env *config.Environment) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(buildURI(env)))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Mongodb Connected to %s", env.MongodbDomain)
	return client.Database(DatabaseName)
}

func buildURI(env *config.Environment) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s", env.MongodbUsername, env.MongodbPassword, env.MongodbDomain, env.MongodbPort)
}
