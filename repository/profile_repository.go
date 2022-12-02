package repository

import (
	"github.com/FTN-TwitterClone/profile/model"
	"golang.org/x/net/context"
)

type ProfileRepository interface {
	SaveUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, username string) (*model.User, error)
	UpdateUser(ctx context.Context, userForm *model.UpdateProfile, authUser *model.AuthUser) error
	DeleteUser(ctx context.Context, username string) error
}
