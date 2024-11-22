package posts

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func (h *Handler) CreateComment(c *gin.Context){
	var context context.Context= c.Request.Context()
	var createCommentRequest posts.CreateCommentRequest
	if err:= c.ShouldBindJSON(&createCommentRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var postIdString string = c.Request.URL.Query().Get("postid") 
	postIdInt64, err := strconv.ParseInt(postIdString, 10, 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var userId int64 = c.GetInt64("userid")
	// var expiry int64 = c.GetInt64("expiry")
	err = h.postsSvc.CreateComment(context,userId,postIdInt64, createCommentRequest)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}
}