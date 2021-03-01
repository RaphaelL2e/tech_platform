package jwtutil

import (
	"fmt"
	"time"

	"gopkg.in/dgrijalva/jwt-go.v3"
)

type JWTConf struct {
	Key      string
	Duration int64
}

type JWTHelper struct {
	Conf JWTConf
}

type CustomClaims struct {
	UserID  string `json:"user_id"`
	IsAdmin bool   `json:"is_admin"`
	jwt.StandardClaims
}

func (h JWTHelper) GenAdminToken(id string, admin bool) (tokenString string, err error) {
	claims := CustomClaims{
		id,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + h.Conf.Duration,
			Issuer:    "login",
		},
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	fmt.Println(h.Conf.Key)
	tokenString, err = token.SignedString([]byte(h.Conf.Key))
	fmt.Println(1,err)
	return
}

func (h JWTHelper) GenToken(id string) (tokenString string, err error) {
	claims := CustomClaims{
		id,
		false,

		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + h.Conf.Duration,
			Issuer:    "login",
		},
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err = token.SignedString([]byte(h.Conf.Key))
	return
}
