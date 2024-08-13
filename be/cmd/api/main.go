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

func run() StatusCode {
	l := log.Default()

	viper, err := initViper()

	if err != nil {
		l.Printf("main: could not unmarshal configuration. %v", err)
		return NotOk
	}

	var cfg app.Config

	err = viper.Unmarshal(&cfg)

	if err != nil {
		l.Printf("main: could not unmarshal configuration. %v", err)
		return NotOk
	}

	app := app.NewApp(l, &cfg)

	l.Printf("main: setup complete")

	g := gin.New()

	g.Use(gin.Recovery())
	g.Use(middlewares.Logger(app))
	g.Use(middlewares.Auth(app))

	handlers.RegisterRoutes(g, app)

	g.SetTrustedProxies(nil)

	l.Printf("main: running")

	g.Run()

	return Ok
}
