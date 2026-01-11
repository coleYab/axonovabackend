package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	MongoDBURI    string
	MongoDBName   string
	Gmail         string
	ReceiverGmail string
	AppPassword   string
	Port          string
}

func NewConfig() *Config {
	if CFG != nil {
		return CFG
	}
	dbUri, err := reqEnv("MONGODB_URI")
	if err != nil {
		log.Fatalln("mongo db is not set in .env")
	}

	CFG = &Config{
		MongoDBURI:    dbUri,
		AppPassword:   getEnvFallback("APP_PASSWORD", ""),
		Gmail:         getEnvFallback("GMAIL", ""),
		ReceiverGmail: getEnvFallback("RECEIVER_GMAIL", ""),
		MongoDBName:   getEnvFallback("MONGODB_NAME", "axonovadb"),
		Port:          getEnvFallback("PORT", ":8080"),
	}
	
	return CFG
}

var CFG *Config = nil

func reqEnv(key string) (string, error) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("env %v does not exist", key)
	}

	return env, nil
}

func getEnvFallback(key string, fallback string) string {
	env, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return env
}
