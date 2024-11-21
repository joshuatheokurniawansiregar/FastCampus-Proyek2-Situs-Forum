package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/configs"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc{
	var secretKey string = configs.GetConfig().Service.SecretJWT
	return func (c *gin.Context){
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		if header == ""{
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userId, userName, err := jwt.ValidateToken(header, secretKey)
		if err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":err.Error(),
			})
			return
		}
		c.Set("userid", userId)
		c.Set("username", userName)
		c.Next()
	}
}