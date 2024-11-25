package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

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

		c.Set("userId", userId)
		c.Set("username", userName)
		c.Next()
		
	}
}

func AuthRefreshMiddleware()gin.HandlerFunc{
	var secretKey string = configs.GetConfig().Service.SecretJWT
	return func(c *gin.Context){
		var header string = c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == ""{
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}
		var(
			userID int64
			username string
			err error
		)
		userID, username, err = jwt.ValidateTokenWithoutExpiry(header, secretKey)
		if err != nil{
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		c.Set("userId", userID)
		c.Set("username", username)
		c.Next()
	}
}

func TokenExpiry()gin.HandlerFunc{
	return func(c *gin.Context){
		var expiry int64 = c.GetInt64("expiry")
		if time.Now().UnixMilli() >= expiry{
			log.Println(expiry)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":errors.New("json token is expired"),
			})
			return
		}
		c.Next()
	}
}