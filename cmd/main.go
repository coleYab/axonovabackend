package main

import (
	"axonova/config"
	server "axonova/internal"
	"axonova/pkg/database"
	"axonova/pkg/mailer"
	"log"
)

func main() {
	cfg := config.NewConfig()
	app := server.NewAppServer()
	db, err := database.NewMongoDB(cfg.MongoDBURI, cfg.MongoDBName)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	gMailer := mailer.NewAppMailer(cfg.Gmail, cfg.AppPassword)
	app.RegisterRoutes(db, gMailer)
	app.Run(cfg.Port)
}
