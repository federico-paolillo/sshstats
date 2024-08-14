package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/federico-paolillo/ssh-attempts/cmd/api/app"
	"github.com/federico-paolillo/ssh-attempts/cmd/api/handlers"
	"github.com/federico-paolillo/ssh-attempts/cmd/api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

func initViper() (*viper.Viper, error) {
	viperInstance := viper.New()

	viperInstance.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viperInstance.SetEnvPrefix("SSHSTATS")

	viperInstance.SetConfigName(".env")
	viperInstance.SetConfigType("dotenv")

	viperInstance.AddConfigPath(".")

	viperInstance.AllowEmptyEnv(false)
	viperInstance.AutomaticEnv()

	viperInstance.SetDefault("server.address", ":65535")

	err := viperInstance.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return viperInstance, nil
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

func run() StatusCode {
	l := log.Default()

	v, err := initViper()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			l.Printf("main: no config file was found. proceeding anyway")
		} else {
			l.Printf("main: could not read configuration. %v", err)
			return NotOk
		}
	}

	var cfg app.Config

	err = v.Unmarshal(&cfg)

	if err != nil {
		l.Printf("main: could not unmarshal configuration. %v", err)
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
