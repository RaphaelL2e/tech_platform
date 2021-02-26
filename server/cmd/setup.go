package main

import (
	"fmt"
	"net/url"
	"tech_platform/server/internal/store/user"
	"tech_platform/server/pkg/jwtutil"

	"github.com/urfave/cli/v2"
)

func setupStore(c *cli.Context) user.Store {
	rdsConfig := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC&time_zone='%s'",
		c.String("db-user"),
		c.String("db-pwd"),
		c.String("db-addr"),
		c.String("db-name"),
		url.QueryEscape("+00:00"))
	return user.New(rdsConfig)
}

func setupJWTHelper(c *cli.Context) jwtutil.JWTHelper {
	return jwtutil.JWTHelper{
		Conf: jwtutil.JWTConf{
			Key:      c.String("jwt-key"),
			Duration: c.Int64("jwt-duration"),
		},
	}
}

