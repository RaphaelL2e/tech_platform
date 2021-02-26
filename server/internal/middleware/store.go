package middleware

import (
	"tech_platform/server/internal/store/user"

	"github.com/gin-gonic/gin"
)
// Store is a middleware function that initializes the store and attaches to
// the context of every http.Request.
func Store(v user.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		user.ToContext(c, v)
		c.Next()
	}
}