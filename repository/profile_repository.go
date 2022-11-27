package repository

import (
	"github.com/FTN-TwitterClone/profile/app_errors"
	"github.com/FTN-TwitterClone/profile/model"
	"golang.org/x/net/context"
)

type ProfileRepository interface {
	SaveUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, username string) (*model.User, *app_errors.AppError)
	UpdateUser(ctx context.Context, userForm *model.UpdateProfile, authUser *model.AuthUser) error
}
