package repository

import (
	"axonova/internal/event/entity"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IEventRepository interface {
	Create(assessment entity.Event) error
	FindByID(id string) (entity.Event, error)
	FindAll() ([]entity.Event, error)
	Update(id string, assessment entity.Event) error
	Delete(id string) error
}

type MongoEventRepository struct {
	collection *mongo.Collection
}

func (m MongoEventRepository) Create(assessment entity.Event) error {
	_, err := m.collection.InsertOne(context.TODO(), toMongoEvent(assessment))
	return err
}

func (m MongoEventRepository) FindByID(id string) (entity.Event, error) {
	var mongoEvent MongoEvent
	if err := m.collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&mongoEvent); err != nil {
		return entity.Event{}, err
	}

	return toEvent(mongoEvent), nil
}

func (m MongoEventRepository) FindAll() ([]entity.Event, error) {
	cur, err := m.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.TODO())
	var events []entity.Event

	for cur.Next(context.TODO()) {
		var mongoEvent MongoEvent
		if err := cur.Decode(&mongoEvent); err != nil {
			return nil, err
		}
		events = append(events, toEvent(mongoEvent))
	}

	return events, nil
}

func (m MongoEventRepository) Update(id string, assessment entity.Event) error {
	oldAssessment, err := m.FindByID(id)
	if err != nil {
		return err
	}

	if oldAssessment.ID != assessment.ID {
		return fmt.Errorf("old assessment id does not match new assessment id")
	}

	m.collection.FindOneAndReplace(context.TODO(), bson.M{"id": id}, toMongoEvent(assessment))
	return nil
}

func (m MongoEventRepository) Delete(id string) error {
	_, err := m.collection.DeleteOne(context.TODO(), bson.M{"id": id})
	return err
}

func NewMongoEventRepository(collection *mongo.Collection) *MongoEventRepository {
	return &MongoEventRepository{
		collection: collection,
	}
}

func toEvent(event MongoEvent) entity.Event {
	return entity.Event{
		ID:           event.ID,
		Title:        event.Title,
		Picture:      event.Picture,
		Description:  event.Description,
		Date:         event.Date,
		StartTime:    event.StartTime,
		DurationMin:  event.DurationMin,
		Price:        event.Price,
		MaxAttendees: event.MaxAttendees,
		IsOnline:     event.IsOnline,
		Platform:     event.Platform,
		MeetingLink:  event.MeetingLink,
		Tags:         event.Tags,
	}
}

func toMongoEvent(event entity.Event) MongoEvent {
	return MongoEvent{
		ID:           event.ID,
		Title:        event.Title,
		Picture:      event.Picture,
		Description:  event.Description,
		Date:         event.Date,
		StartTime:    event.StartTime,
		DurationMin:  event.DurationMin,
		Price:        event.Price,
		MaxAttendees: event.MaxAttendees,
		IsOnline:     event.IsOnline,
		Platform:     event.Platform,
		MeetingLink:  event.MeetingLink,
		Tags:         event.Tags,
	}
}

type MongoEvent struct {
	ID           string    `bson:"id"`
	Title        string    `bson:"title"`
	Picture      string    `bson:"picture"`
	Description  string    `bson:"description"`
	Date         time.Time `bson:"date"`
	StartTime    string    `bson:"startTime"`
	DurationMin  int       `bson:"minDuration"`
	Price        int       `bson:"price"`
	MaxAttendees int       `bson:"maxAttendees"`
	IsOnline     bool      `bson:"isOnline"`
	Platform     string    `bson:"platform"`
	MeetingLink  string    `bson:"meetingLink"`
	Tags         []string  `bson:"tags"`
}
