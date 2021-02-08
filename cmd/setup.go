package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net/url"
	"tech_platform/internal/pkg/jwtauth"
	"tech_platform/internal/user/store"
)

func setupStore(c *cli.Context) store.Store {
	rdsConfig := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC&time_zone='%s'",
		c.String("db-user"),
		c.String("db-pwd"),
		c.String("db-addr"),
		c.String("db-name"),
		url.QueryEscape("+00:00"))
	return store.New(rdsConfig)
}

func setupJWTHelper(c *cli.Context) *jwtauth.JWTHelper {
	return &jwtauth.JWTHelper{Conf: jwtauth.JWTConf{
		Key:      c.String("jwt-key"),
		Duration: 24 * 60 * 60, // 1 day
	}}
}
