package technology

import (
	"github.com/gin-gonic/gin"
	"tech_platform/server/internal/server/technology"
	"tech_platform/server/pkg/jwtutil"
)

var (
	srv *technology.Handler
)

func TechnologyRouter0(router *gin.RouterGroup,helper jwtutil.JWTHelper){
	g0 :=router.Group("/technology")
	{
		g0.GET("/get/:id",)
	}

	srv = technology.NewHandler()
}

func TechnologyRouter1(router *gin.RouterGroup){
	g0 :=router.Group("/technology")
	{
		g0.POST("/add",)
		g0.PUT("/update",)
		g0.DELETE("/delete",)
	}
}


