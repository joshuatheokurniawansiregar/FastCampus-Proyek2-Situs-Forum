package posts

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func (s *service) CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest)error{
	var hashtags string = strings.Join(req.PostHashtags, ",") 
	var data posts.PostModel = posts.PostModel{
		UserId: userId,
		PostTitle: req.PostTitle,
		PostContent: req.PostContent,
		PostHashtags: hashtags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: strconv.FormatInt(userId, 10),
		UpdatedBy: strconv.FormatInt(userId, 10),
	}
	err := s.postsRepo.CreatePost(ctx, data)
	if err != nil{
		return err
	}
	return nil
}