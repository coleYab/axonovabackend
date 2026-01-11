package repository

import (
	"axonova/internal/assesment/entity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAssessmentRepository interface {
	Create(assessment entity.Assessment) error
	FindByID(id string) (entity.Assessment, error)
	FindAll() ([]entity.Assessment, error)
	Update(id string, assessment entity.Assessment) error
	Delete(id string) error
}

type MongoAssessmentRepository struct {
	collection *mongo.Collection
}

func NewMongoAssessmentRepository(collection *mongo.Collection) *MongoAssessmentRepository {
	return &MongoAssessmentRepository{
		collection: collection,
	}
}

func toMongoAssessment(assessment entity.Assessment) MongoAssessment {
	return MongoAssessment{
		ID:                  assessment.ID,
		Name:                assessment.Name,
		Email:               assessment.Email,
		Company:             assessment.Company,
		Answers:             assessment.Answers,
		TotalScore:          assessment.TotalScore,
		RecommendationTitle: assessment.RecommendationTitle,
		AnsweredCount:       assessment.AnsweredCount,
		TotalQuestions:      assessment.TotalQuestions,
	}
}

func (m MongoAssessmentRepository) Create(assessment entity.Assessment) error {
	mongoAssessment := toMongoAssessment(assessment)
	_, err := m.collection.InsertOne(context.Background(), mongoAssessment)
	return err
}

func getContext() context.Context {
	return context.TODO()
}

func toAssessment(mAssessment MongoAssessment) entity.Assessment {
	return entity.Assessment{
		ID:                  mAssessment.ID,
		Name:                mAssessment.Name,
		Email:               mAssessment.Email,
		Company:             mAssessment.Company,
		Answers:             mAssessment.Answers,
		TotalScore:          mAssessment.TotalScore,
		RecommendationTitle: mAssessment.RecommendationTitle,
		AnsweredCount:       mAssessment.AnsweredCount,
		TotalQuestions:      mAssessment.TotalQuestions,
	}
}

func (m MongoAssessmentRepository) FindByID(id string) (entity.Assessment, error) {
	ctx := getContext()
	mongoAssessment := MongoAssessment{}
	if err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&mongoAssessment); err != nil {
		return entity.Assessment{}, err
	}

	return toAssessment(mongoAssessment), nil
}

func (m MongoAssessmentRepository) FindAll() ([]entity.Assessment, error) {
	ctx := getContext()
	var assessments []entity.Assessment
	cur, err := m.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		mongoEntity := MongoAssessment{}
		if err := cur.Decode(&mongoEntity); err != nil {
			return nil, err
		}
		assessments = append(assessments, toAssessment(mongoEntity))
	}

	return assessments, nil
}

func (m MongoAssessmentRepository) Update(id string, assessment entity.Assessment) error {
	//TODO implement me
	panic("implement me")
}

func (m MongoAssessmentRepository) Delete(id string) error {
	_, err := m.collection.DeleteOne(getContext(), bson.M{"_id": id})
	return err
}

type MongoAssessment struct {
	ID      string `bson:"_id,omitempty"`
	Name    string `bson:"name"`
	Email   string `bson:"email"`
	Company string `bson:"company,omitempty"`

	Answers    map[string]int `bson:"answers"`
	TotalScore int            `bson:"totalScore"`

	RecommendationTitle string `bson:"recommendationTitle"`

	AnsweredCount  int `bson:"answeredCount"`
	TotalQuestions int `bson:"totalQuestions"`
}
