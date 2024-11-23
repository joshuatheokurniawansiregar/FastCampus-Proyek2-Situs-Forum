package posts

import (
	"context"
	"errors"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func (s *service) GetPostById(ctx context.Context, postId int64)(posts.Post,error){
	var(
		post posts.Post
		err error
	)

	post, err = s.postsRepo.GetPostById(ctx, postId)
	if err != nil{
		return post, errors.New("get_post_error: "+err.Error())
	}

	post.LikesTotal, err = s.postsRepo.CountUserLikesByPostId(ctx, postId)
	if err != nil{
		return post, errors.New("get_total_likes_error: "+err.Error())
	}

	post.Comments,err= s.postsRepo.GetCommentsByPostIdID(ctx, postId)
	if err != nil{
		return post, errors.New("get_comments_error: "+err.Error())
	}
	return post, nil
}