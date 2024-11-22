package posts

import (
	"context"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/configs"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

type postsRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context,  model posts.CommentModel) error
	CreateUserLike(ctx context.Context, userLikesModel posts.UserLikesModel)error
	UpdateUserLike(ctx context.Context, userLikesModel posts.UserLikesModel)error
	GetAllPosts(ctx context.Context, limit, offset int64)(posts.GetAllPostsResponse, error)
}

type service struct{
	postsRepo postsRepository
	cfg *configs.Config
}

func NewService(postsRepo postsRepository, cfg *configs.Config)*service{
	return &service{
		postsRepo: postsRepo,
		cfg: cfg,
	}
}