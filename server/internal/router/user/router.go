package user

import (
	"github.com/gin-gonic/gin"
	"tech_platform/server/internal/server/user"
	"tech_platform/server/pkg/jwtutil"
)

var (
	srv *user.Handler
)

func UserRouter0(router *gin.RouterGroup, helper jwtutil.JWTHelper) {
	g0 := router.Group("/user")
	{
		g0.POST("/register", register)
		g0.POST("/login", login)
		g0.GET("/:userId", getUserinfo)
	}

	srv = user.NewHandler(helper)
}

func UserRouter1(router *gin.RouterGroup) {
	g1 := router.Group("/user")
	{
		g1.POST("/updateUserinfo", updateUserInfo)
	}
}
