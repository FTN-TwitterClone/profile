package saga

import (
	"encoding/json"
	"fmt"
	"github.com/FTN-TwitterClone/profile/repository"
	"github.com/FTN-TwitterClone/profile/tracing"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
	"os"
)

type RegisterUserHandler struct {
	tracer            trace.Tracer
	conn              *nats.Conn
	profileRepository repository.ProfileRepository
}

func NewRegisterUserHandler(tracer trace.Tracer, profileRepository repository.ProfileRepository) (*RegisterUserHandler, error) {
	natsHost := os.Getenv("NATS_HOST")
	natsPort := os.Getenv("NATS_PORT")

	url := fmt.Sprintf("nats://%s:%s", natsHost, natsPort)

	connection, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	h := &RegisterUserHandler{
		tracer:            tracer,
		conn:              connection,
		profileRepository: profileRepository,
	}

	_, err = connection.Subscribe(REGISTER_COMMAND, h.handleCommand)
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (h RegisterUserHandler) handleCommand(msg *nats.Msg) {
	remoteCtx, err := tracing.GetNATSParentContext(msg)
	if err != nil {

	}

	ctx, span := otel.Tracer("profile").Start(trace.ContextWithRemoteSpanContext(context.Background(), remoteCtx), msg.Subject)
	defer span.End()

	var c RegisterUserCommand

	err = json.Unmarshal(msg.Data, &c)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return
	}

	switch c.Command {
	case SaveProfile:
		h.handleSaveProfile(ctx, c.User)
	case RollbackProfile:
		h.handleRollbackProfile(ctx, c.User)
	}
}

func (h RegisterUserHandler) handleSaveProfile(ctx context.Context, user NewUser) {
	_, span := h.tracer.Start(ctx, "RegisterUserHandler.handleSaveProfile")
	defer span.End()
}

func (h RegisterUserHandler) handleRollbackProfile(ctx context.Context, user NewUser) {
	_, span := h.tracer.Start(ctx, "RegisterUserHandler.handleRollbackProfile")
	defer span.End()
}
