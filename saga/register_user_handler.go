package saga

import (
	"fmt"
	"github.com/FTN-TwitterClone/profile/repository"
	"github.com/nats-io/nats.go"
	"os"
)

type RegisterUserHandler struct {
	conn              *nats.EncodedConn
	profileRepository repository.ProfileRepository
}

func NewRegisterUserHandler(profileRepository repository.ProfileRepository) (*RegisterUserHandler, error) {
	natsHost := os.Getenv("NATS_HOST")
	natsPort := os.Getenv("NATS_PORT")

	url := fmt.Sprintf("nats://%s:%s", natsHost, natsPort)

	connection, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	encConn, err := nats.NewEncodedConn(connection, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}

	h := &RegisterUserHandler{
		conn:              encConn,
		profileRepository: profileRepository,
	}

	_, err = encConn.Subscribe(REGISTER_COMMAND, h.handleCommand)
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (h RegisterUserHandler) handleCommand(c RegisterUserCommand) {
	switch c.Command {
	case SaveProfile:
		h.handleSaveProfile(c.User)
	case RollbackProfile:
		h.handleRollbackProfile(c.User)
	}
}

func (h RegisterUserHandler) handleSaveProfile(user NewUser) {

}

func (h RegisterUserHandler) handleRollbackProfile(user NewUser) {

}
