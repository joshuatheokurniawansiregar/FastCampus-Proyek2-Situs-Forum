package memberships

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/memberships"
)

// "github.com/gin-gonic/gin"

func (h *Handler) Login(c *gin.Context){
	var ctx context.Context= c.Request.Context()

	var request memberships.LoginRequest
	
	if err := c.ShouldBindJSON(&request); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	accessToken, refreshToken, err := h.membershipsSvc.Login(ctx, request)

	var response memberships.LoginResponse = memberships.LoginResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}