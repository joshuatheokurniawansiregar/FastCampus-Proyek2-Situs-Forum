package posts

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/posts"
)

func (h *Handler) GetAllPosts(c *gin.Context){
	var context context.Context = c.Request.Context()
	
	var(
		pageSize string
		pageSizeInt64 int64

		pageNumber string
		pageNumberInt64 int64

		err error
	)

	pageSize = c.Request.URL.Query().Get("pagesize")
	pageSizeInt64, err = strconv.ParseInt(pageSize, 10, 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	pageNumber = c.Request.URL.Query().Get("pagenumber")
	pageNumberInt64, err = strconv.ParseInt(pageNumber, 10, 64)
	
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	var(
		getAllPostsResponse posts.GetAllPostsResponse
	)

	getAllPostsResponse, err = h.postsSvc.GetAllPosts(context, pageSizeInt64, pageNumberInt64)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"allPostsResponse": getAllPostsResponse,
	})
}