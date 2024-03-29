package main

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tech_platform/server/internal/middleware"
	"tech_platform/server/internal/router"
	"tech_platform/server/pkg/ossutil"
	"time"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "db-addr",
		Usage:   "database address",
		EnvVars: []string{"DB_ADDR"},
	},
	&cli.StringFlag{
		Name:    "db-user",
		Usage:   "database user",
		EnvVars: []string{"DB_USER"},
	},
	&cli.StringFlag{
		Name:    "db-pwd",
		Usage:   "database pwd",
		EnvVars: []string{"DB_PWD"},
	},
	&cli.StringFlag{
		Name:    "db-name",
		Usage:   "database name",
		EnvVars: []string{"DB_NAME"},
	},
	&cli.StringFlag{
		Name:    "jwt-key",
		Usage:   "jwt-key",
		EnvVars: []string{"JWT_KEY"},
		Value:   "this is key",
	},
	&cli.Int64Flag{
		Name:    "jwt-duration",
		Usage:   "jwt-duration",
		EnvVars: []string{"JWT_DURATION"},
		Value:   24 * 60 * 60,
	},
}

func server(c *cli.Context) (err error) {
	store := setupStore(c)
	jwtHelper := setupJWTHelper(c)
	ossHelper :=ossutil.NewClient()
	handler := router.Setup(
		c,
		jwtHelper,
		ossHelper,
		middleware.Store(store),
		middleware.Cors(),
	)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if serr := srv.Shutdown(ctx); serr != nil {
		logrus.Fatal("Server forced to shutdown:", serr)
	}

	logrus.Println("Server exiting")

	return
}
