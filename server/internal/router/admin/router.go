package admin

import (
	"github.com/gin-gonic/gin"
	"tech_platform/server/internal/server/admin"
	"tech_platform/server/pkg/jwtutil"
)

var (
	srv *admin.Handler
)

func AdminRouter0(router *gin.RouterGroup,helper jwtutil.JWTHelper){
	g0 :=router.Group("/admin")
	{
		g0.POST("/login",adminLogin)
	}

	srv = admin.NewHandler(helper)
}

func AdminRouter1(router *gin.RouterGroup){
	g1 := router.Group("/admin")
	{
		g1.POST("/add",adminAdd)
		g1.POST("/list",adminList)
	}
}


