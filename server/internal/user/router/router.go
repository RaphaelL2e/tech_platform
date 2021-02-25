package router

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"tech_platform/server/internal/middleware"
	"tech_platform/server/internal/user/server"
	"tech_platform/server/pkg/jwtutil"
)

var (
	srv *server.Handler
)

func init() {

}

func Setup(c *cli.Context, helper1 jwtutil.JWTHelper, middlewares ...gin.HandlerFunc) *gin.Engine {
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
	srv = server.NewHandler(helper1)

	g0 := router.Group("/apiv11/user")
	{
		g0.POST("/register", register)
		g0.POST("/login", login)
		g0.Use(middleware.JWTAuth(c.String("jwt-key"))).POST("/updateUserinfo", updateUserInfo)
		g0.GET("/:userId", getUserinfo)
	}

	return router
}
