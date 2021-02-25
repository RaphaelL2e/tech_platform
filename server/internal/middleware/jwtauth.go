package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"tech_platform/server/internal/pkg/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuth is a middleware function that check jwt token.
func JWTAuth(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := parseRequest(c.Request)
		if err != nil {
			resp := response.CreateByErrorCodeMessage(response.ForbiddenCode)
			c.JSON(http.StatusUnauthorized, resp)
			c.Abort()
			return
		}

		token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
			// validate the alg
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(key), nil
		})

		if err != nil {
			fmt.Println(2)
			resp := response.CreateByErrorCodeMessage(response.ForbiddenCode)
			c.JSON(http.StatusUnauthorized, resp)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
			c.Set("is_admin", claims["is_admin"])
		} else {
			fmt.Println(3)
			resp := response.CreateByErrorCodeMessage(response.ForbiddenCode)
			c.JSON(http.StatusUnauthorized, resp)
			c.Abort()
			return
		}
	}
}

func parseRequest(r *http.Request) (*string, error) {
	// first we attempt to get the token from the
	// authorization header.
	var token = r.Header.Get("Authorization")
	if len(token) != 0 {
		return &token, nil
	}

	// then we attempt to get the token from the
	// access_token url query parameter
	token = r.FormValue("access_token")
	if len(token) != 0 {
		return &token, nil
	}

	// and finally we attempt to get the token from
	// the user session cookie
	cookie, err := r.Cookie("user_sess")
	if err == nil {
		token = cookie.Value
		if len(token) != 0 {
			return &token, nil
		}
	}

	return nil, errors.New("not found")
}
