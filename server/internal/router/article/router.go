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
		g0.POST("/list",list)
	}
	srv = article.NewHandler()
}

func Articlerouter1(router *gin.RouterGroup){
	g1 :=router.Group("/article")
	{
		g1.POST("/add",addArticle)
	}
}






