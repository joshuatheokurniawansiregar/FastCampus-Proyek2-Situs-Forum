package posts

import "time"

type (
	CreateCommentRequest struct {
		CommentContent string `json:"commentContent"`
	}
)

type (
	CommentModel struct {
		ID        int64  `db:"id"`
		UserId    int64  `db:"user_id"`
		PostId    int64  `db:"post_id"`
		Comments  string `db:"comments"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string `db:"created_by"`
		UpdatedBy string `db:"updated_by"`
	}
)