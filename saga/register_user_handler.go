package saga

import (
	"encoding/json"
	"fmt"
	"github.com/FTN-TwitterClone/profile/model"
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

	ctx, span := otel.Tracer("profile").Start(trace.ContextWithRemoteSpanContext(context.Background(), remoteCtx), "RegisterUserHandler.handleCommand")
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
	handlerCtx, span := h.tracer.Start(ctx, "RegisterUserHandler.handleSaveProfile")
	defer span.End()

	var u model.User

	if user.Role == "ROLE_USER" {
		u = model.User{
			Username:    user.Username,
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Town:        user.Town,
			Gender:      user.Gender,
			YearOfBirth: user.YearOfBirth,
			Private:     true,
		}
	} else {
		u = model.User{
			Username:    user.Username,
			Email:       user.Email,
			CompanyName: user.CompanyName,
			Website:     user.Website,
			Private:     false,
		}
	}

	err := h.profileRepository.SaveUser(handlerCtx, &u)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		h.sendReply(handlerCtx, RegisterUserReply{
			Reply: ProfileFail,
			User:  user,
		})

		return
	}

	h.sendReply(handlerCtx, RegisterUserReply{
		Reply: ProfileSuccess,
		User:  user,
	})
}

func (h RegisterUserHandler) handleRollbackProfile(ctx context.Context, user NewUser) {
	handlerCtx, span := h.tracer.Start(ctx, "RegisterUserHandler.handleRollbackProfile")
	defer span.End()

	err := h.profileRepository.DeleteUser(handlerCtx, user.Username)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return
	}

	h.sendReply(handlerCtx, RegisterUserReply{
		Reply: ProfileRollback,
		User:  user,
	})
}

func (h RegisterUserHandler) sendReply(ctx context.Context, r RegisterUserReply) {
	_, span := h.tracer.Start(ctx, "RegisterUserHandler.sendReply")
	defer span.End()

	headers := nats.Header{}
	headers.Set(tracing.TRACE_ID, span.SpanContext().TraceID().String())
	headers.Set(tracing.SPAN_ID, span.SpanContext().SpanID().String())

	data, err := json.Marshal(r)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return
	}

	msg := nats.Msg{
		Subject: REGISTER_REPLY,
		Header:  headers,
		Data:    data,
	}

	err = h.conn.PublishMsg(&msg)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	}
}
