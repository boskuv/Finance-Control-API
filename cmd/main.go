package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	fc "github.com/boskuv/Finance-Control-API"
	"github.com/boskuv/Finance-Control-API/pkg/handler"

	//"github.com/boskuv/Finance-Control-API/pkg/service"

	"github.com/sirupsen/logrus"
)

// @title Finance Control API
// @version 1.0
// @description API Server for Finance Control Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// if err := godotenv.Load(); err != nil {
	// 	logrus.Fatalf("error loading env variables: %s", err.Error())
	// }

	handlers := new(handler.Handler)

	srv := new(fc.Server)
	if err := srv.Run( "8008", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
