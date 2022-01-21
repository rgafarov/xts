package app

import (
	"context"

	"github.com/rgafarov/xts/util/logger"
)

type Server interface {
	Run(app *App) error
	Shutdown()
}

type App struct {
	Name   string
	Config string
	Server Server
	Logger *logger.Logger
}

func New(name string, configFile string, server Server) *App {
	return &App{Logger: logger.NewConsole(true),
		Server: server}
}

func (a *App) Init() error {
	a.Logger.Info().Msgf("%s starting", a.Name)
	return nil
}

func (a *App) Run() error {
	a.Logger.Info().Msgf("%s running", a.Name)
	a.Server.Run(a)
	return nil
}

func (a *App) Shutdown(ctx context.Context) {
	a.Server.Shutdown()
	a.Logger.Info().Msg("shutdown")
}
