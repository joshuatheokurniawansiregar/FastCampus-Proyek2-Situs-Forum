package memberships

import (
	"context"
	"errors"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/memberships"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest)(string, error){
	userModelByEmail , err  := s.membershipRepo.GetUserForLogin(ctx, req.Email)
	if err != nil{	
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}
	if userModelByEmail == nil{
		return "", errors.New("email does not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userModelByEmail.Password), []byte(req.Password))
	if err != nil{
		return "", errors.New("password is invalid")
	}

	token, err := jwt.CreateToken(userModelByEmail.Id, userModelByEmail.UserName, s.cfg.Service.SecretJWT)
	if err !=nil{
		return "", nil
	}
	return token, nil
}