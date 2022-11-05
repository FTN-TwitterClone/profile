package service

import (
	"context"
	"github.com/FTN-TwitterClone/profile/proto/profile"
	"github.com/FTN-TwitterClone/profile/repository"
	"github.com/golang/protobuf/ptypes/empty"
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
	_, span := s.tracer.Start(ctx, "gRPCProfileService.RegisterUser")
	defer span.End()

	println("yay")

	return new(empty.Empty), nil
}

func (s gRPCProfileService) RegisterBusinessUser(ctx context.Context, businessUser *profile.BusinessUser) (*empty.Empty, error) {
	_, span := s.tracer.Start(ctx, "gRPCProfileService.RegisterBusinessUser")
	defer span.End()

	println("yay")

	return new(empty.Empty), nil
}
