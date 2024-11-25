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

	RefreshTokenRequest struct{
		Token string `json:"token"`
	}
)

type(
	LoginResponse struct{
		AccessToken string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}
	RefreshResponse struct{
		AccessToken string `json:"accessToken"`
	}
)

type (
	UsersModel struct {
		ID        int64    `db:"id"`
		Email     string    `db:"email"`
		Password  string    `db:"password"`
		UserName  string    `db:"user_name"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string `db:"created_by"`
		UpdatedBy string `db:"updated_by"`
	}

	RefreshTokenModel struct{
		ID int64 `db:"id"`
		UserID int64 `db:"user_id"`
		RefreshToken string `db:"refresh_token"`
		ExpiredAt time.Time `db:"refresh_token"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string `db:"created_by"`
		UpdatedBy string `db:"updated_by"`
	}
)