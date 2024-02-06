package main

import (
	"context"
	"fmt"
	initial "github.com/Max425/WB-Tech-level-0/cmd"
	"github.com/Max425/WB-Tech-level-0/pkg/api"
	"github.com/Max425/WB-Tech-level-0/pkg/api/handler"
	"github.com/Max425/WB-Tech-level-0/pkg/repository"
	"github.com/Max425/WB-Tech-level-0/pkg/service"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title WB API
// @version 1.0
// @description API Server for Umlaut Application

// @host localhost:8000
// @BasePath /
func main() {
	initial.InitConfig()
	logger, err := initial.InitLogger()
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Starting server...")
	defer logger.Sync()

	ctx := context.Background()

	db, err := initial.InitPostgres(ctx)
	if err != nil {
		logger.Error("initialize Postgres",
			zap.String("Error", fmt.Sprintf("failed to initialize Postgres: %s", err.Error())))
	}

	redis, err := initial.InitRedis()
	if err != nil {
		logger.Error("initialize redisDb",
			zap.String("Error", fmt.Sprintf("failed to initialize redisDb: %s", err.Error())))
	}
	defer redis.Close()

	repos := repository.NewRepository(db, redis, logger)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, logger)

	srv := new(api.Server)

	//TODO: LOAD CACHE

	go func() {
		if err = srv.Serve(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logger.Error("running http server",
				zap.String("Error", fmt.Sprintf("error occured while running http server: %s", err.Error())))
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("TodoApp Shutting Down")

	if err = srv.Shutdown(context.Background()); err != nil {
		logger.Error("error occured on server shutting down: %s",
			zap.Error(err))
	}
}
