package service

import (
	"github.com/FTN-TwitterClone/profile/proto/profile"
	"github.com/FTN-TwitterClone/profile/repository"
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
