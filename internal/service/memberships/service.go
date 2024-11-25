package memberships

import (
	"context"
	"time"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/configs"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/memberships"
)

type membershipsRepository interface {
	//Get user returns UsersModel pointer struct by email, UsersModel pointer struct by user_name, and error
	GetUser(ctx context.Context, email string, userName string, userID int64)(*memberships.UsersModel, error)
	GetUserForSignin(ctx context.Context, email, userName string)(*memberships.UsersModel, *memberships.UsersModel, error)
	CreateUser(ctx context.Context, model memberships.UsersModel)error
	InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel)error
	GetRefreshToken(ctx context.Context, userId int64, now time.Time)(*memberships.RefreshTokenModel, error)
}

type service struct{
	membershipRepo membershipsRepository // interface
	cfg *configs.Config
}

func NewService(membershipsRepo membershipsRepository, cfg *configs.Config) *service{
	return &service{
		membershipRepo: membershipsRepo,
		cfg: cfg,
	}
}