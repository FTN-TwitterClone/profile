package service

import (
	"github.com/FTN-TwitterClone/profile/app_errors"
	"github.com/FTN-TwitterClone/profile/model"
	"github.com/FTN-TwitterClone/profile/repository"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
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

func (s *ProfileService) GetUser(ctx context.Context, username string) (*model.User, *app_errors.AppError) {
	serviceCtx, span := s.tracer.Start(ctx, "ProfileService.GetUser")
	defer span.End()

	user, err := s.profileRepository.GetUser(serviceCtx, username)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, &app_errors.AppError{500, ""}
	}

	return user, nil
}
