package mongo

import (
	"go.opentelemetry.io/otel/trace"
)

type MongoProfileRepository struct {
	tracer trace.Tracer
}

func NewMongoProfileRepository(tracer trace.Tracer) (*MongoProfileRepository, error) {
	//db := os.Getenv("DB")
	//dbport := os.Getenv("DBPORT")

	//mongo logic

	car := MongoProfileRepository{
		tracer,
	}

	return &car, nil
}
