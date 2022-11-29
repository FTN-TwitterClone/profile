package service

import (
	"github.com/FTN-TwitterClone/grpc-stubs/proto/social_graph"
	"github.com/FTN-TwitterClone/profile/app_errors"
	"github.com/FTN-TwitterClone/profile/model"
	"github.com/FTN-TwitterClone/profile/repository"
	"github.com/FTN-TwitterClone/profile/tls"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
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
func (s *ProfileService) SaveUser(ctx context.Context, user *model.User) *app_errors.AppError {
	serviceCtx, span := s.tracer.Start(ctx, "ProfileService.SaveUser")
	defer span.End()

	err := s.profileRepository.SaveUser(serviceCtx, user)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return &app_errors.AppError{500, "User not saved."}
	}

	return nil
}

func (s *ProfileService) UpdateUser(ctx context.Context, userForm *model.UpdateProfile, authUser *model.AuthUser) *app_errors.AppError {
	serviceCtx, span := s.tracer.Start(ctx, "ProfileService.UpdateUser")
	defer span.End()

	err := s.profileRepository.UpdateUser(serviceCtx, userForm, authUser)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return &app_errors.AppError{500, "User not saved."}
	}

	conn, err := getgRPCConnection("social-graph:9001")
	defer conn.Close()
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return &app_errors.AppError{500, ""}
	}

	privacy := social_graph.SocialGraphUpdatedUser{
		Private: userForm.Private,
	}

	socialGraphService := social_graph.NewSocialGraphServiceClient(conn)
	serviceCtx = metadata.AppendToOutgoingContext(serviceCtx, "authUsername", authUser.Username)

	_, err = socialGraphService.SocialGraphUpdateUser(serviceCtx, &privacy)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return &app_errors.AppError{500, ""}
	}

	return nil
}

func getgRPCConnection(address string) (*grpc.ClientConn, error) {
	creds := credentials.NewTLS(tls.GetgRPCClientTLSConfig())

	conn, err := grpc.DialContext(
		context.Background(),
		address,
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)

	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}

	return conn, err
}
