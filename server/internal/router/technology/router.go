package technology

import (
	"github.com/gin-gonic/gin"
	"tech_platform/server/internal/server/technology"
)

var (
	srv *technology.Handler
)

func TechnologyRouter0(router *gin.RouterGroup){
	g0 :=router.Group("/technology")
	{
		g0.GET("/get/:id",getTechnology)
	}

	srv = technology.NewHandler()
}

func TechnologyRouter1(router *gin.RouterGroup){
	g0 :=router.Group("/technology")
	{
		g0.POST("/add",addTechnology)
		g0.PUT("/update",updateTechnology)
		g0.DELETE("/delete",deleteTechnology)
	}
}


