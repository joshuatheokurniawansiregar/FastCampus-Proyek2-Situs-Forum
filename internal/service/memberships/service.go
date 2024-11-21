package memberships

import (
	"context"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/configs"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/memberships"
)

type membershipsRepository interface {
	//Get user returns UsersModel pointer struct by email, UsersModel pointer struct by user_name, and error
	GetUser(ctx context.Context, email string, username string) (*memberships.UsersModel, *memberships.UsersModel, error)
	GetUserForLogin(ctx context.Context, email string)(*memberships.UsersModel, error)
	CreateUser(ctx context.Context, model memberships.UsersModel)error
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