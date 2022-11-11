package service

import (
	"context"
	"github.com/FTN-TwitterClone/profile/app_errors"
	"github.com/FTN-TwitterClone/profile/model"
	"github.com/FTN-TwitterClone/profile/proto/profile"
	"github.com/FTN-TwitterClone/profile/repository"
	"github.com/golang/protobuf/ptypes/empty"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type gRPCProfileService struct {
	profile.UnimplementedProfileServiceServer
	tracer            trace.Tracer
	profileRepository repository.ProfileRepository
}

func NewgRPCProfileService(tracer trace.Tracer, profileRepository repository.ProfileRepository) *gRPCProfileService {
	return &gRPCProfileService{
		tracer:            tracer,
		profileRepository: profileRepository,
	}
}

func (s gRPCProfileService) RegisterUser(ctx context.Context, user *profile.User) (*empty.Empty, error) {
	serviceCtx, span := s.tracer.Start(ctx, "gRPCProfileService.RegisterUser")
	defer span.End()
	u := model.ProfileUser{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Town:      user.Town,
		Gender:    user.Gender,
		//Age:	user.Age
	}
	err := s.profileRepository.SaveUser(serviceCtx, &u)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, &app_errors.AppError{500, ""}
	}
	return new(empty.Empty), nil
}

func (s gRPCProfileService) RegisterBusinessUser(ctx context.Context, businessUser *profile.BusinessUser) (*empty.Empty, error) {
	_, span := s.tracer.Start(ctx, "gRPCProfileService.RegisterBusinessUser")
	defer span.End()

	println("yay")

	return new(empty.Empty), nil
}
