package memberships

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/memberships"
)

func (h *Handler) Refresh(c *gin.Context){
	var context context.Context = c.Request.Context()

	var request memberships.RefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userId := c.GetInt64("userId")
	accessToken, err := h.membershipsSvc.ValidateRefreshToken(context, userId, request)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, memberships.RefreshResponse{
		AccessToken: accessToken,
	})
}