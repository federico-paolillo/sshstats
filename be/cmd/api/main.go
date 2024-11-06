package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/federico-paolillo/ssh-attempts/cmd/api/app"
	"github.com/federico-paolillo/ssh-attempts/cmd/api/handlers"
	"github.com/federico-paolillo/ssh-attempts/cmd/api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/sherifabdlnaby/configuro"
)

type StatusCode = int

var (
	Ok    StatusCode = 0
	NotOk            = 1
)

func main() {
	executionResult := run()
	os.Exit(executionResult)
}

func initServer(app *app.App) *http.Server {
	g := gin.New()

	g.Use(gin.Recovery())
	g.Use(middlewares.Logger(app))
	g.Use(middlewares.Auth(app))

	handlers.RegisterRoutes(g, app)

	g.SetTrustedProxies(nil)

	s := &http.Server{
		Addr:         app.Cfg.Server.Address,
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 1 * time.Second,
		Handler:      g,
	}

	return s
}

func initConfig(logger *log.Logger) (app.Config, error) {
	var cfg app.Config

	c, err := configuro.NewConfig(
		configuro.WithoutExpandEnvVars(),
		configuro.WithLoadFromEnvVars("SSHSTATS"),
		configuro.WithLoadFromConfigFile("config.yml", false),
		configuro.WithEnvConfigPathOverload("SSHSTATS_CONFIG_DIR"),
	)

	if err != nil {
		return cfg, fmt.Errorf(
			"main: could not setup configuro. %v",
			err,
		)
	}

	err = c.Load(&cfg)

	if err != nil {
		return cfg, fmt.Errorf(
			"main: could not load configuration. %v",
			err,
		)
	}

	azFunctionsPort := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")

	if azFunctionsPort != "" {
		logger.Printf(
			"we are running as az function. port is %s",
			azFunctionsPort,
		)

		cfg.Server.Address = net.JoinHostPort("", azFunctionsPort)
	}

	return cfg, nil
}

func run() StatusCode {
	l := log.Default()

	cfg, err := initConfig(l)

	if err != nil {
		l.Fatalf(
			"main: failed to init config. %v",
			err,
		)

		return NotOk
	}

	app := app.NewApp(l, &cfg)

	l.Printf("main: setup complete")

	s := initServer(app)

	l.Printf("main: listening on %s", s.Addr)

	err = s.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		return Ok
	}

	l.Printf("main: server closed unexpectedly. %v", err)

	return NotOk
}
