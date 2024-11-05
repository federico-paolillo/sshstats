package main

import (
	"errors"
	"log"
	"net"
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
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("dotenv")

	v.AddConfigPath(".")
	v.AddConfigPath("/etc/sshstats")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix("SSHSTATS")
	v.AllowEmptyEnv(false)
	v.AutomaticEnv()

	v.SetDefault("server.address", ":65535")

	v.SetDefault("auth.headerkey", "")
	v.SetDefault("auth.headervalue", "")

	v.SetDefault("loki.username", "")
	v.SetDefault("loki.password", "")
	v.SetDefault("loki.endpoint", "")

	v.BindEnv("server.address")

	v.BindEnv("auth.headerkey")
	v.BindEnv("auth.headervalue")

	v.BindEnv("loki.username")
	v.BindEnv("loki.password")
	v.BindEnv("loki.endpoint")

	err := v.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return v, nil
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

	cfg := &app.Config{
		Loki: app.LokiSettings{
			Endpoint: v.GetString("loki.endpoint"),
			Username: v.GetString("loki.username"),
			Password: v.GetString("loki.password"),
		},
		Server: app.ServerSettings{
			Address: v.GetString("server.address"),
		},
		Auth: app.AuthSettings{
			HeaderKey:   v.GetString("auth.headerkey"),
			HeaderValue: v.GetString("auth.headervalue"),
		},
	}

	azFunctionsPort := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")

	if azFunctionsPort != "" {
		l.Printf(
			"we are running as az function. port is %s",
			azFunctionsPort,
		)

		cfg.Server.Address = net.JoinHostPort("", azFunctionsPort)
	}

	app := app.NewApp(l, cfg)

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
