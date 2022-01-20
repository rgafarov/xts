package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rgafarov/xts/cmd/authserver/app"
	"github.com/rgafarov/xts/cmd/authserver/server"
)

func main() {

	const appName = "auth"
	configFile := ""

	// configuring signals for interruption
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)

	// creating new instance of application
	a := app.New(appName, configFile, &server.AuthServer{})

	// initialise the application
	if a.Init() != nil {
		return
	}

	// start main application loop in go routine
	go a.Run()

	// waiting for shutdown
	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shutdown appplication
	a.Shutdown(ctx)
}
