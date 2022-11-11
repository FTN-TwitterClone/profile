package repository

import (
	"github.com/FTN-TwitterClone/profile/model"
	"golang.org/x/net/context"
)

type ProfileRepository interface {
	SaveUser(ctx context.Context, user *model.ProfileUser) error
}
