package app

import (
	"context"

	"github.com/rgafarov/xts/cmd/authserver/server"
	"github.com/rgafarov/xts/util/logger"
)

type App struct {
	Name   string
	Config string
	Server server.Server
	Logger *logger.Logger
}

func New(name string, configFile string, server server.Server) *App {
	return &App{Logger: logger.NewConsole(true),
		Server: server}
}

func (a *App) Init() error {
	a.Logger.Info().Msg("started")
	return nil
}

func (a *App) Run() error {
	a.Logger.Info().Msg("running")
	a.Server.Run()
	return nil
}

func (a *App) Shutdown(ctx context.Context) {
	a.Server.Shutdown()
	a.Logger.Info().Msg("shutdown")
}
