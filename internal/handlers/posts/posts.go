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
	CreateUserLike(ctx context.Context, userId int64, postId int64, request posts.UserLikeRequest)error
	UpdateUserLike(ctx context.Context, userId int64, postId int64, request posts.UserLikeRequest)error
	GetAllPosts(ctx context.Context, pageSize, pageNumber int64)(posts.GetAllPostsResponse, error)
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
	route.GET("/allposts",handler.GetAllPosts)
	route.POST("/create", handler.CreatePost)
	route.POST("/comments/create", handler.CreateComment)
	route.POST("user_likes/create", handler.CreateUserLike)
	route.POST("user_likes/update", handler.UpdateUserLike)
}

