package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/LKarlon/test_task/pkg/handler"
	"github.com/LKarlon/test_task/pkg/server"
	"github.com/LKarlon/test_task/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {

	port := "8080"
	service := service.NewService()
	handlers := handler.NewHandler(service)
	srv := new(server.Server)
	go func() {
		if err := srv.Run(port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("Converter Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Converter Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
