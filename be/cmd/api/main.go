package main

import (
	"fmt"
	"log"
	"os"

	"github.com/federico-paolillo/ssh-attempts/cmd/api/app"
	"github.com/federico-paolillo/ssh-attempts/cmd/api/handlers"
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

	viperInstance.SetEnvPrefix("RSFLLO")
	viperInstance.AllowEmptyEnv(false)
	viperInstance.AutomaticEnv()

	err := viperInstance.ReadInConfig()

	if err != nil {
		return nil, fmt.Errorf("main: could not read configuration. %v", err)
	}

	return viperInstance, nil
}

func run() StatusCode {
	log := log.Default()

	viper, err := initViper()

	if err != nil {
		log.Printf("main: could not init configuration. %w", err)
		return NotOk
	}

	var cfg app.Config

	err = viper.Unmarshal(&cfg)

	if err != nil {
		log.Printf("main: could not unmarshal configuration. %w", err)
		return NotOk
	}

	app := app.NewApp(log, &cfg)

	log.Printf("main: up setup")

	gin := gin.Default()

	handlers.RegisterRoutes(gin, app)

	return Ok
}
