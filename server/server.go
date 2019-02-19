package server

import (
	"github.com/labstack/echo"
	em "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"context"
	"time"
	"boardbots/server/routes/makegame"
	"boardbots/manager"
	"boardbots/server/middleware"
	"boardbots/server/routes/joingame"
	"boardbots/server/routes/newuser"
	"boardbots/server/persistence"
	"boardbots/server/transport"
	"boardbots/server/routes/signin"
	"boardbots/server/routes/getgames"
)

func StartEchoServer() {
	server := echo.New()
	server.Logger.SetLevel(log.DEBUG)
	server.Use(em.Logger())
	transport.EchoErrorHandler(server)
	userPortal := persistence.NewMemoryPortal()
	gameManager := manager.NewMemoryGameManager()

	newuser.ApplyRoute(server, &userPortal)
	signin.ApplyRoute(server, &userPortal)
	api := server.Group("/api", middleware.ContextHander)

	h := makegame.Handler{
		GameManager: gameManager,
	}
	api.POST("/makegame", h.MakeGame)
	gamesApi := api.Group("/g")
	joinGameHandler := joingame.Handler{}
	gamesApi.POST("/join", joinGameHandler.JoinGame)
	getgames.ApplyRoute(gamesApi, gameManager)


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