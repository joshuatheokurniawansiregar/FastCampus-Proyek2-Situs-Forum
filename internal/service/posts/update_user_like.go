package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func (s *service) UpdateUserLike(ctx context.Context, userId int64, postId int64, request posts.UserLikeRequest)error{
	var userLikesModel posts.UserLikesModel = posts.UserLikesModel{
		UserId: userId,
		PostId: postId,
		IsLiked: request.IsLIked,
		UpdatedAt: time.Now(),
		UpdatedBy: strconv.FormatInt(userId, 10),
	}
	var err error = s.postsRepo.UpdateUserLike(ctx, userLikesModel)
	if err != nil{
		return err
	}
	return nil
}