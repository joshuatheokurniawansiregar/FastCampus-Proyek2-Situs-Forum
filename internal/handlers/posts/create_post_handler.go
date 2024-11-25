package posts

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func (handler *Handler) CreatePost(c *gin.Context) {
	var context context.Context = c.Request.Context()
	var request posts.CreatePostRequest

	if err := c.ShouldBindJSON(&request); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	var userId int64 = c.GetInt64("userId")
	var err error = handler.postsSvc.CreatePost(context,userId,request)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}