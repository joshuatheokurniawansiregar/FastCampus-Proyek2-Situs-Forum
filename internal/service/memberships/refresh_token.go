package memberships

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/memberships"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/pkg/jwt"
	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest)(string, error){
	var(
		existingRefreshToken *memberships.RefreshTokenModel
		err error
		user *memberships.UsersModel
		token string
	)
	existingRefreshToken, err = s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil{
		return "", err
	}
	fmt.Println("userID: ", userID)
	fmt.Println("refreshToken", request.Token)
	if existingRefreshToken == nil{
		return "", errors.New("refresh token has expired")
	}
	if existingRefreshToken.RefreshToken != request.Token{
		return "", errors.New("refresh token is invalid")
	}
	user, err = s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil{
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}
	if user == nil{
		return "", errors.New("user not exist")
	}
	token, err = jwt.CreateToken(user.ID, user.UserName,s.cfg.Service.SecretJWT)
	if err != nil{
		return "", err
	}
	return token, nil
}