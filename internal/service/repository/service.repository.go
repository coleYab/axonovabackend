package repository

import (
	"axonova/internal/service/entity"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IServiceRepository interface {
	CreateService(assessment entity.Service) error
	CreateContact(assessment entity.Contact) error
	FindByIDService(id string) (entity.Service, error)
	FindByIDContact(id string) (entity.Contact, error)
	FindAllService() ([]entity.Service, error)
	FindAllContact() ([]entity.Contact, error)
}

type MongoServiceRepository struct {
	collection *mongo.Collection
}

func NewMongoServiceRepository(collection *mongo.Collection) *MongoServiceRepository {
	return &MongoServiceRepository{
		collection: collection,
	}
}

type MongoContact struct {
	ID      string `bson:"id,omitempty"` // MongoDB default ID field
	Name    string `bson:"name"`
	Email   string `bson:"email"`
	Message string `bson:"message"`
}

type MongoService struct {
	ID               string   `bson:"id,omitempty"`
	Name             string   `bson:"name"`
	Email            string   `bson:"email"`
	Message          string   `bson:"message"`
	Phone            string   `bson:"phone,omitempty"`
	Service          string   `bson:"service"`
	PreferredDate    string   `bson:"preferred_date"`
	RequestedModules []string `bson:"requested_modules"`
}

func toMongoContact(c entity.Contact) MongoContact {
	return MongoContact{
		ID:      c.ID,
		Name:    c.Name,
		Email:   c.Email,
		Message: c.Message,
	}
}

func fromMongoContact(m MongoContact) entity.Contact {
	return entity.Contact{
		ID:      m.ID,
		Name:    m.Name,
		Email:   m.Email,
		Message: m.Message,
	}
}

func toMongoService(s entity.Service) MongoService {
	return MongoService{
		ID:               s.ID,
		Name:             s.Name,
		Email:            s.Email,
		Message:          s.Message,
		Phone:            s.Phone,
		Service:          s.Service,
		PreferredDate:    s.PreferredDate,
		RequestedModules: s.RequestedModules,
	}
}

func fromMongoService(m MongoService) entity.Service {
	return entity.Service{
		ID:               m.ID,
		Name:             m.Name,
		Email:            m.Email,
		Message:          m.Message,
		Phone:            m.Phone,
		Service:          m.Service,
		PreferredDate:    m.PreferredDate,
		RequestedModules: m.RequestedModules,
	}
}

// --- Implementation Methods ---

func (r *MongoServiceRepository) CreateService(service entity.Service) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mService := toMongoService(service)
	_, err := r.collection.InsertOne(ctx, mService)
	return err
}

func (r *MongoServiceRepository) CreateContact(contact entity.Contact) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mContact := toMongoContact(contact)
	_, err := r.collection.InsertOne(ctx, mContact)
	return err
}

func (r *MongoServiceRepository) FindByIDService(id string) (entity.Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var m MongoService
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&m)
	if err != nil {
		return entity.Service{}, err
	}
	return fromMongoService(m), nil
}

func (r *MongoServiceRepository) FindByIDContact(id string) (entity.Contact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var m MongoContact
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&m)
	if err != nil {
		return entity.Contact{}, err
	}
	return fromMongoContact(m), nil
}

func (r *MongoServiceRepository) FindAllService() ([]entity.Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filter for documents that have a 'service' field
	cursor, err := r.collection.Find(ctx, bson.M{"service": bson.M{"$exists": true}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []entity.Service
	for cursor.Next(ctx) {
		var m MongoService
		if err := cursor.Decode(&m); err != nil {
			return nil, err
		}
		results = append(results, fromMongoService(m))
	}
	return results, nil
}

func (r *MongoServiceRepository) FindAllContact() ([]entity.Contact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filter for documents that do NOT have a 'service' field
	cursor, err := r.collection.Find(ctx, bson.M{"service": bson.M{"$exists": false}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []entity.Contact
	for cursor.Next(ctx) {
		var m MongoContact
		if err := cursor.Decode(&m); err != nil {
			return nil, err
		}
		results = append(results, fromMongoContact(m))
	}
	return results, nil
}
