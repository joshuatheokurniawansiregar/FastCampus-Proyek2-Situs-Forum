package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest)error{
	// fmt.Println("email: "+req.Email)
	// fmt.Println("username: "+req.Username)
	byEmail, byUserName, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)
	if err != nil{
		return err
	}

	if byEmail != nil && byUserName != nil{
		return errors.New("email and username are already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	var now time.Time = time.Now()

	var user_model memberships.UsersModel= memberships.UsersModel{
		Email     : req.Email,
		Password  : string(pass),
		UserName  :req.Username,
		CreatedAt :now,
		UpdatedAt :now,
		CreatedBy :req.Email,
		UpdatedBy :req.Email,
	}

	return s.membershipRepo.CreateUser(ctx, user_model)
	
}