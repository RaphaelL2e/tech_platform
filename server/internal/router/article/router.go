package article

import (
	"github.com/gin-gonic/gin"
	"tech_platform/server/internal/server/article"
)

var (
	srv *article.Handler
)

func ArticleRouter0(router *gin.RouterGroup)  {
	g0 :=router.Group("/article")
	{
		g0.GET("/:id",getArticleById)
	}
	srv = article.NewHandler()
}


