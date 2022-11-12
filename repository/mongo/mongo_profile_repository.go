package mongo

import (
	"context"
	"fmt"
	"github.com/FTN-TwitterClone/profile/app_errors"
	"github.com/FTN-TwitterClone/profile/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"os"
)

type MongoProfileRepository struct {
	tracer trace.Tracer
	cli    *mongo.Client
}

func NewMongoProfileRepository(tracer trace.Tracer) (*MongoProfileRepository, error) {

	db := os.Getenv("DB")
	dbport := os.Getenv("DBPORT")

	//mongo logic
	host := fmt.Sprintf("%s:%s", db, dbport)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(`mongodb://`+host))
	if err != nil {
		panic(err)
	}
	client.Database("twitterCloneDB").Collection("users")

	car := MongoProfileRepository{
		tracer,
		client,
	}

	return &car, nil
}

func (r *MongoProfileRepository) SaveUser(ctx context.Context, user *model.User) error {
	_, span := r.tracer.Start(ctx, "MongoProfileRepository.SaveUser")
	defer span.End()

	usersCollection := r.cli.Database("twitterCloneDB").Collection("users")

	_, err := usersCollection.InsertOne(ctx, user)

	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	return nil
}

func (r *MongoProfileRepository) GetUser(ctx context.Context, username string) (*model.User, *app_errors.AppError) {
	_, span := r.tracer.Start(ctx, "MongoProfileRepository.GetUser")
	defer span.End()

	usersCollection := r.cli.Database("twitterCloneDB").Collection("users")
	result := usersCollection.FindOne(ctx, bson.M{"username": username})

	var elem *model.User

	err := result.Decode(elem)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, &app_errors.AppError{500, ""}
	}

	return elem, nil
}
