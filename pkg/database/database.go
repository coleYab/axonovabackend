package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoDB(mongoURI, dbname string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	// ping
	if err := client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(context.Background())
		return nil, fmt.Errorf("failed to ping mongodb %v", err)
	}

	// get the database
	db := client.Database(dbname)

	return &MongoDB{
		client:   client,
		database: db,
	}, nil
}

func (db *MongoDB) GetCollection(collectionName string) *mongo.Collection {
	return db.database.Collection(collectionName)
}

func (db *MongoDB) Close() error {
	return db.client.Disconnect(context.Background())
}
