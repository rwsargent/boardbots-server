package server

import (
	"github.com/labstack/echo"
	em "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"context"
	"time"
	"boardbots-server/server/routes/makegame"
	"boardbots-server/manager"
	"boardbots-server/server/middleware"
	"boardbots-server/server/routes/newuser"
	"boardbots-server/server/persistence"
	"boardbots-server/server/transport"
	"boardbots-server/server/routes/signin"
	"boardbots-server/server/routes/getgames"
	"boardbots-server/server/routes/joingame"
)

const ApiPrefix = "/api/v0"

func StartEchoServer() {
	server := echo.New()
	server.Logger.SetLevel(log.DEBUG)
	server.Use(em.Logger())
	transport.EchoErrorHandler(server)
	userPortal := persistence.NewMemoryPortal()
	gameManager := manager.NewMemoryGameManager()

	newuser.ApplyRoute(server, userPortal)
	signin.ApplyRoute(server, userPortal)

	api := server.Group(ApiPrefix, middleware.ContextHander)
	api.Use(em.BasicAuthWithConfig(middleware.GetBasicAuthenticator(userPortal)))
	makegame.ApplyRoute(api, gameManager)

	gamesApi := api.Group("/g", middleware.Authenticator(gameManager))
	getgames.ApplyRoute(gamesApi, gameManager)
	joingame.ApplyRoute(gamesApi, gameManager)


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