package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func (s *service) CreateUserLike(ctx context.Context, userId int64, postId int64, request posts.UserLikeRequest)error{
	var userLikesModel posts.UserLikesModel = posts.UserLikesModel{
		UserId: userId,
		PostId: postId,
		IsLiked: request.IsLIked,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: strconv.FormatInt(userId, 10),
		UpdatedBy: strconv.FormatInt(userId, 10),
	}
	var err error = s.postsRepo.CreateUserLike(ctx, userLikesModel)
	
	if err != nil{
		return err
	}
	return nil
}