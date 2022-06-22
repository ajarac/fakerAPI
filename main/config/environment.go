package config

import (
	"os"
)

type Environment struct {
	MongodbDomain   string
	MongodbPort     string
	MongodbUsername string
	MongodbPassword string
	RapidAPIKey     string
	Port            string
}

func GetEnvironment() *Environment {
	return &Environment{
		MongodbDomain:   getEnvOrDefault("MONGODB_DOMAIN", "localhost"),
		MongodbPort:     getEnvOrDefault("MONGODB_PORT", "27017"),
		MongodbUsername: getEnvOrDefault("MONGODB_USERNAME", "root"),
		MongodbPassword: getEnvOrDefault("MONGODB_PASSWORD", "password"),
		RapidAPIKey:     getEnvOrDefault("RAPID_API_KEY", ""),
		Port:            getEnvOrDefault("PORT", "3000"),
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
