package middleware

import (
	"github.com/gin-gonic/gin"
	"tech_platform/internal/pkg/jwtauth"
)

func VerifyToken(helper *jwtauth.JWTHelper) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("token")
		if len(tokenString) == 0 {

			c.Abort()
			return
		}

		if err := helper.VerifyToken(tokenString); err != nil {

			c.Abort()
			return
		}
	}
}
