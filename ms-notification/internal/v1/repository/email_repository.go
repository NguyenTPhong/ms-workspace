package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"ms-workspace/ms-notification/internal/v1/repository/model"
)

type EmailRepository interface {
	SaveEmail(ctx context.Context, email *model.Email) error
}

type emailRepository struct {
	mongoDb *mongo.Database
}

func NewEmailRepository(database *mongo.Database) EmailRepository {
	repository := &emailRepository{
		mongoDb: database,
	}
	return repository
}

func (e *emailRepository) SaveEmail(ctx context.Context, email *model.Email) error {
	collection := e.mongoDb.Collection("sent_emails")
	_, err := collection.InsertOne(ctx, email)
	if err != nil {
		return err
	}
	return nil
}
