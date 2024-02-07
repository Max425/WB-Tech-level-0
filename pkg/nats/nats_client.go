package nats

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/service"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Nats struct {
	ctx      context.Context
	sc       stan.Conn
	logger   *zap.Logger
	services *service.Service
}

func StartNatsClient(ctx context.Context, logger *zap.Logger, services *service.Service) {
	sc, err := stan.Connect(
		viper.GetString("nats.clusterID"),
		viper.GetString("nats.clientID"),
		stan.NatsURL(viper.GetString("nats.url")),
	)
	if err != nil {
		logger.Error("Error stan Connect", zap.Error(err))
		return
	}
	natsWorker := NewNats(ctx, sc, logger, services)
	go natsWorker.StartClient()
}

func NewNats(ctx context.Context, sc stan.Conn, logger *zap.Logger, services *service.Service) *Nats {
	return &Nats{
		ctx:      ctx,
		sc:       sc,
		logger:   logger,
		services: services,
	}
}

func (n *Nats) StartClient() {
	sub, err := n.sc.Subscribe(viper.GetString("nats.subject"), func(msg *stan.Msg) {
		_, err := n.services.Order.CreateOrder(n.ctx, msg.Data)
		if err != nil {
			n.logger.Error("error nats client: %s", zap.Error(err))
			return
		}
	}, stan.StartWithLastReceived())
	if err != nil {
		n.logger.Error("error nats client: %s", zap.Error(err))
	}
	defer sub.Unsubscribe()

	n.logger.Info("Subscriber connected and listening for messages...")

	select {}
}
