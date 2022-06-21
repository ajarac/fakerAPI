package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Environment struct {
	MongodbDomain string
}

func GetEnvironment() *Environment {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln("Error loading environments")
	}
	return &Environment{
		MongodbDomain: getEnvOrDefault("MONGODB_DOMAIN", "mongodb://root:password@localhost:27017"),
	}
}

func getEnvOrDefault(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	} else {
		return val
	}
}
