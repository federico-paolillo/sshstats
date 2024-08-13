package main

import (
	"fmt"
	"log"
	"os"

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

	viperInstance.SetConfigName("config")
	viperInstance.SetConfigType("json")

	viperInstance.AddConfigPath(".")
	viperInstance.AddConfigPath("/etc/sshstats")

	viperInstance.SetEnvPrefix("SSHSTATS")
	viperInstance.AllowEmptyEnv(false)
	viperInstance.AutomaticEnv()

	err := viperInstance.ReadInConfig()

	if err != nil {
		return nil, fmt.Errorf("main: could not read config. file. %w", err)
	}

	return viperInstance, nil
}

func initGin(app *app.App) *gin.Engine {
	g := gin.New()

	g.Use(gin.Recovery())
	g.Use(middlewares.Logger(app))
	g.Use(middlewares.Auth(app))

	handlers.RegisterRoutes(g, app)

	g.SetTrustedProxies(nil)

	return g
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

	g := initGin(app)

	l.Printf("main: running")

	g.Run()

	return Ok
}
