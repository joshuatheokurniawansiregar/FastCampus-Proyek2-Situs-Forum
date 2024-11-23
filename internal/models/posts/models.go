package posts

import "time"

type (
	CreatePostRequest struct {
		PostTitle    string   `json:"postTitle"`
		PostContent  string   `json:"postContent"`
		PostHashtags []string `json:"postHashTags"`
	}
)

type (
	PostModel struct {
		ID           int64  `db:"id"`
		UserId       int64  `db:"user_id"`
		PostTitle    string `db:"post_title"`
		PostContent  string `db:"post_content"`
		PostHashtags string `db:"post_hashtags"`
		CreatedAt    time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string `db:"created_by"`
		UpdatedBy string `db:"updated_by"`
	}
)

type(
	GetAllPostsResponse struct{
		Data []Post `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	Post struct{
		ID int64 `json:"id"`
		UserId int64 `json:"userId"`
		Username string `json:"username"`
		PostTitle string `json:"postTitle"`
		PostContent string `json:"postContent"`
		PostHashTags []string `json:"postHashTags"`
		CreatedAt time.Time `json:"createdAt"`
		LikesTotal int64 `json:"likesTotal"`
		Comments []Comment `json:"comments"`
	}

	Pagination struct{
		Limit int64 `json:"limit"`
		Offset int64 `json:"offset"`
	}

	GetPostResponse struct{
		PostDetail Post `json:"postDetail"`
		LikeCount int `json:"likeCount"`
		Comments []Comment `json:"comments"`
	}

	Comment struct{
		ID int64 `json:"id"`
		UserId int64 `json:"user_id"`
		Username string `json:"username"`
		CommentContent string `json:"commentContent"`
	}
)