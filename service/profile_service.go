package service

import (
	"github.com/FTN-TwitterClone/profile/repository"
	"go.opentelemetry.io/otel/trace"
)

type ProfileService struct {
	tracer            trace.Tracer
	profileRepository repository.ProfileRepository
}

func NewProfileService(tracer trace.Tracer, profileRepository repository.ProfileRepository) *ProfileService {
	return &ProfileService{
		tracer,
		profileRepository,
	}
}
