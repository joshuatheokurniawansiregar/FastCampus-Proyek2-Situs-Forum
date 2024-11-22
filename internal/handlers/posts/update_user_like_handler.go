package posts

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func (h *Handler) UpdateUserLike(c *gin.Context){
	var context context.Context = c.Request.Context()
	var userLikeRequest posts.UserLikeRequest
	if err := c.ShouldBindJSON(&userLikeRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var(
		postIdString string
		postIdInt64 int64
		err error
	)

	postIdString = c.Request.URL.Query().Get("postid")
	postIdInt64, err = strconv.ParseInt(postIdString, 10, 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	var userId int64 = c.GetInt64("userid")
	err = h.postsSvc.UpdateUserLike(context, userId, postIdInt64, userLikeRequest)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}