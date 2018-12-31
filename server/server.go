package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"context"
	"time"
	"boardbots/server/routes/makegame"
)

func StartEchoServer() {
	server := echo.New()
	server.Logger.SetLevel(log.DEBUG)

	// Apply Routes
	h := makegame.Handler{}
	server.POST("/makegame", h.MakeGame)

	go func() {
		if err := server.Start(":8080"); err != nil {
			server.Logger.Infof("error on init, shutting down server. %v\n", err)
		}
	}()

	shutdownGracefully(server)
}

func shutdownGracefully(server *echo.Echo) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}