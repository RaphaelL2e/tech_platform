package router

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"tech_platform/server/internal/middleware"
	"tech_platform/server/internal/router/admin"
	"tech_platform/server/internal/router/article"
	"tech_platform/server/internal/router/user"
	"tech_platform/server/internal/router/technology"
	"tech_platform/server/pkg/jwtutil"
)

func Setup(c *cli.Context,jwtHelper jwtutil.JWTHelper, middlewares ...gin.HandlerFunc) *gin.Engine {
	gin.DisableConsoleColor()

	// Creates a router without any middleware by default
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	router.Use(middlewares...)

	PublicGroup := router.Group("/api/v1")
	{
		user.UserRouter0(PublicGroup,jwtHelper)
		admin.AdminRouter0(PublicGroup,jwtHelper)
		technology.TechnologyRouter0(PublicGroup)
		article.ArticleRouter0(PublicGroup)
	}

	PrivateGroup := router.Group("/api/v1")
	PrivateGroup.Use(middleware.JWTAuth(c.String("jwt-key")))
	{
		user.UserRouter1(PrivateGroup)
		admin.AdminRouter1(PrivateGroup)
		technology.TechnologyRouter1(PrivateGroup)
		article.Articlerouter1(PrivateGroup)
	}

	return router
}
