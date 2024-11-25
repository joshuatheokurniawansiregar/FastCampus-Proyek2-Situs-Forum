package memberships

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/memberships"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/pkg/jwt"
	tokenUtil "github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/pkg/token"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error){
	userModelByEmail, err  := s.membershipRepo.GetUser(ctx, req.Email, "", 0)
	if err != nil{	
		log.Error().Err(err).Msg("failed to get user")
		return "","", err
	}
	if userModelByEmail == nil{
		return "","", errors.New("email does not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userModelByEmail.Password), []byte(req.Password))
	if err != nil{
		return "" , "", errors.New("password is invalid")
	}

	token, err := jwt.CreateToken(userModelByEmail.ID, userModelByEmail.UserName, s.cfg.Service.SecretJWT)
	if err !=nil{
		return "","", nil
	}

	var existingRefreshToken *memberships.RefreshTokenModel
	existingRefreshToken, err = s.membershipRepo.GetRefreshToken(ctx, userModelByEmail.ID, time.Now())
	
	if err != nil{
		log.Error().Err(err).Msg("error get latest refresh token from database")
		return "", "", err
	}
	
	if existingRefreshToken != nil{
		return token, existingRefreshToken.RefreshToken, nil
	}

	var refreshToken string = tokenUtil.GenerateRefreshToken()
	if refreshToken == ""{
		return token, "", errors.New("failed to generate refresh token")
	}

	err = s.membershipRepo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserID:  userModelByEmail.ID,
		RefreshToken: refreshToken,
		ExpiredAt: time.Now().Add(10 * 24 * time.Hour),
		CreatedAt:  time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: strconv.FormatInt(userModelByEmail.ID, 10),
		UpdatedBy: strconv.FormatInt(userModelByEmail.ID, 10),
	})

	if err != nil{
		return token, refreshToken, err
	}
	return token,refreshToken, nil
}