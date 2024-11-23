package posts

import (
	"context"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func (s *service) GetAllPosts(ctx context.Context, pageSize, pageNumber int64)(posts.GetAllPostsResponse, error){
	var(
		limit int64
		offset int64

		allPostsResponse posts.GetAllPostsResponse
		err error
	)
	limit = pageSize
	offset = pageSize * (pageNumber - 1)
	allPostsResponse, err = s.postsRepo.GetAllPosts(ctx, limit, offset)
	if err != nil{
		return allPostsResponse, err
	}
	return allPostsResponse, nil
}