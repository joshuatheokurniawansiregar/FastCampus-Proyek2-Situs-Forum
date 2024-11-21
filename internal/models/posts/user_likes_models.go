package posts

import "time"

type (
	UserLikeRequest struct {
		IsLIked bool `json:"isLiked"`
	}
)

type (
	UserLikesModel struct {
		Id        int64 `db:"id"`
		UserId    int64 `db:"user_id"`
		PostId    int64 `db:"post_id"`
		IsLiked   bool  `db:"is_liked"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string `db:"created_by"`
		UpdatedBy string `db:"updated_by"`
	}
)
