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

	v1 := router.Group("/api/v1/user")
	{
		v1.POST("/register", register)
		v1.POST("/login", login)
		v1.Use(middleware.JWTAuth(c.String("jwt-key"))).POST("/updateUserinfo", updateUserInfo)
		v1.GET("/:userId",getUserinfo)
	}

	return router
}
