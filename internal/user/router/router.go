package router

import (
	"tech_platform/internal/pkg/jwtauth"
	"tech_platform/internal/user/server"

	"github.com/gin-gonic/gin"
)

var (
	srv *server.Handler

	jwt *jwtauth.JWTHelper
)

func init() {

}

func Setup(helper *jwtauth.JWTHelper, middleware ...gin.HandlerFunc) *gin.Engine {
	gin.DisableConsoleColor()

	// Creates a router without any middleware by default
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	srv = server.NewHandler()
	jwt = helper

	return router
}
