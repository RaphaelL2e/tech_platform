package file

import (
	"github.com/gin-gonic/gin"
	"tech_platform/server/internal/server/file"
	"tech_platform/server/pkg/ossutil"
)


var (
	srv *file.Handler
)

func FileRouter0(router *gin.RouterGroup){
}

func FileRouter1(router *gin.RouterGroup,helper ossutil.OSSHelper){
	g1 :=router.Group("/file")
	{
		g1.POST("/upload",uploadfile)

	}
	srv = file.NewHandler(helper)
}