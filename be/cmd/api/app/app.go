package app

import (
	"log"

	"github.com/federico-paolillo/ssh-attempts/internal/caching"
	"github.com/federico-paolillo/ssh-attempts/internal/loki"
	"github.com/federico-paolillo/ssh-attempts/pkg/stats"
)

type App struct {
	Cfg      *Config
	Log      *log.Logger
	Provider stats.Provider
}

func NewApp(
	log *log.Logger,
	cfg *Config,
) *App {
	app := &App{
		Log: log,
		Cfg: cfg,
	}

	initStatProvider(app)

	return app
}

func initStatProvider(app *App) {
	app.Provider = caching.NewProvider(
		app.Log,
		loki.NewProvider(
			app.Log,
			loki.NewLogcliConnector(
				loki.NewClient(&loki.Config{
					User:     app.Cfg.Loki.Username,
					Password: app.Cfg.Loki.Password,
					Endpoint: app.Cfg.Loki.Endpoint,
				}),
			),
		),
	)
}
