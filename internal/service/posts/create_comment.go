package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func (s *service) CreateComment(ctx context.Context, userId int64, postId int64, request posts.CreateCommentRequest)error{
	var commentModel posts.CommentModel = posts.CommentModel{
		UserId: userId,
		PostId: postId,
		Comments: request.CommentContent,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: strconv.FormatInt(userId,10),
		UpdatedBy: strconv.FormatInt(userId,10),
	}
	err:= s.postsRepo.CreateComment(ctx, commentModel)
	if err != nil{
		return err
	}
	return nil
}