package memberships

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/models/memberships"
)

type membershipsService interface{
	SignUp(ctx context.Context, req memberships.SignUpRequest)error
	Login(ctx context.Context, req memberships.LoginRequest)(string,error)
}

type Handler struct {
	*gin.Engine

	membershipsSvc membershipsService
}

func NewHandler(api *gin.Engine, memberhipSvc membershipsService) *Handler {
	return &Handler{
		Engine: api,
		membershipsSvc: memberhipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	var route *gin.RouterGroup = h.Group("memberships")
	route.GET("/ping", h.Ping)
	route.POST("/signup", h.SignUp)
	route.POST("/login", h.Login)
}
