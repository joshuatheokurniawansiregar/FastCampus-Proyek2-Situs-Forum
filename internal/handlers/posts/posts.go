package posts

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/middleware"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

type postsService interface {
	CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, userId int64, postId int64, request posts.CreateCommentRequest)error
}

type Handler struct{
	*gin.Engine
	postsSvc postsService
}

func NewHandler(api *gin.Engine, postsSvc postsService)*Handler{
	return &Handler{
		Engine: api,
		postsSvc: postsSvc,
	}
}

func(handler *Handler) RegisterRoute(){
	route:= handler.Group("posts")
	route.Use(middleware.AuthMiddleware())
	route.POST("/create", handler.CreatePost)
	route.POST("/comments/create", handler.CreateComment)
}

