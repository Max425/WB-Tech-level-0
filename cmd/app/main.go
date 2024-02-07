package main

import (
	"context"
	"fmt"
	"github.com/Max425/WB-Tech-level-0/cmd"
	"github.com/Max425/WB-Tech-level-0/cmd/nats"
	"github.com/Max425/WB-Tech-level-0/pkg/api"
	"github.com/Max425/WB-Tech-level-0/pkg/api/handler"
	"github.com/Max425/WB-Tech-level-0/pkg/repository"
	"github.com/Max425/WB-Tech-level-0/pkg/service"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title WB API
// @version 1.0
// @description API Server for WB-level-0 Application

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
	services := service.NewService(repos, logger)
	handlers := handler.NewHandler(services, logger)
	srv := new(api.Server)

	go func() {
		if err = srv.Serve(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logger.Error("error occurred on server shutting down", zap.Error(err))
			os.Exit(1)
		}
	}()
	err = services.Order.LoadOrdersToCache(ctx)
	if err != nil {
		logger.Error("error load cache", zap.Error(err))
	}
	go startNats(logger)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("WB Shutting Down")

	if err = srv.Shutdown(context.Background()); err != nil {
		logger.Error("error occurred on server shutting down: %s", zap.Error(err))
	}
}

func startNats(logger *zap.Logger) {
	sc, err := stan.Connect(
		viper.GetString("nats.clusterID"),
		viper.GetString("nats.clientID"),
		stan.NatsURL(viper.GetString("nats.url")),
	)
	if err != nil {
		logger.Error("Error stan Connect", zap.Error(err))
		return
	}
	natsWorker := nats.NewNats(sc, logger)
	go natsWorker.StartClient()
	go natsWorker.StartServer()
}
