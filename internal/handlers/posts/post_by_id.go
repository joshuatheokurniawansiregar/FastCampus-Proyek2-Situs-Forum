package posts

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func (h *Handler) GetPostById(c *gin.Context){
	var(
		context context.Context
		postIdString string
		postId int64
		post posts.Post
		err error
	)
	context = c.Request.Context()
	postIdString = c.Request.URL.Query().Get("postid")
	postId,err = strconv.ParseInt(postIdString, 10, 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	post, err = h.postsSvc.GetPostById(context, postId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}