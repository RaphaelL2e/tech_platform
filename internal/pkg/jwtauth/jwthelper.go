package jwtauth

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type JWTConf struct {
	Key      string
	Duration int64
}

type JWTHelper struct {
	Conf JWTConf
}

func (h *JWTHelper) GenToken(begin int64) (tokenString string, err error) {
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: begin + h.Conf.Duration,
		Issuer:    "liyafei",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err = token.SignedString([]byte(h.Conf.Key))
	return
}

func (h *JWTHelper) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(h.Conf.Key), nil
	})

	if err != nil {
		return err
	} else if !token.Valid {
		return errors.New("invalid token")
	} else {
		return nil
	}
}
