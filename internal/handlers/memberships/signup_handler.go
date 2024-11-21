package memberships

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/memberships"
)

func (h *Handler) SignUp(c *gin.Context) {
	var context context.Context= c.Request.Context()

	var request memberships.SignUpRequest
	if err:= c.ShouldBindJSON(&request); err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err:= h.membershipsSvc.SignUp(context, request)
	if err!= nil{
		email_error_message := fmt.Sprintf("Error 1062 (23000): Duplicate entry '%s' for key 'users.email'", request.Email)
		user_name_error_messge := fmt.Sprintf("Error 1062 (23000): Duplicate entry '%s' for key 'users.user_name'", request.Username)
		if err.Error() == email_error_message{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "email is already exists",
			})
		}
		if err.Error() == user_name_error_messge{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "username is already exists",
			})
		}

		if err.Error() != email_error_message && err.Error() != user_name_error_messge{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	}

	c.JSON(http.StatusCreated,nil)
}