package admin

import (
	"github.com/gin-gonic/gin"
	"tech_platform/server/internal/model/user"
	"tech_platform/server/pkg/jwtutil"
)

type Handler struct {
	jwtHelper jwtutil.JWTHelper
}

func NewHandler(helper jwtutil.JWTHelper) *Handler {
	return &Handler{jwtHelper: helper}
}

func (h Handler) AdminLogin(c *gin.Context, req user.LoginRequest) (user.LoginResponse, error) {

	return user.LoginResponse{}, nil
}

