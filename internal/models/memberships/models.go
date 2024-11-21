package memberships

import "time"

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	}

	LoginRequest struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}
)

type(
	LoginResponse struct{
		AccessToken string `json:"access_token"`
		
	}
)

type (
	UsersModel struct {
		Id        int64    `db:"id"`
		Email     string    `db:"email"`
		Password  string    `db:"password"`
		UserName  string    `db:"user_name"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string `db:"created_by"`
		UpdatedBy string `db:"updated_by"`
	}
)