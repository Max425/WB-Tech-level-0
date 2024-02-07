package nats

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/Max425/WB-Tech-level-0/pkg/constants"
	"github.com/Max425/WB-Tech-level-0/pkg/model/dto"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"time"
)

type Nats struct {
	sc     stan.Conn
	logger *zap.Logger
}

func NewNats(sc stan.Conn, logger *zap.Logger) *Nats {
	return &Nats{
		sc:     sc,
		logger: logger,
	}
}

func (n *Nats) StartClient() {
	sub, err := n.sc.Subscribe(viper.GetString("nats.subject"), func(msg *stan.Msg) {
		log.Printf("Received a message: %s\n", string(msg.Data))
	}, stan.StartWithLastReceived())
	if err != nil {
		n.logger.Error("error nats client: %s", zap.Error(err))
	}
	defer sub.Unsubscribe()

	n.logger.Info("Subscriber connected and listening for messages...")

	select {}
}

func (n *Nats) StartServer() {
	n.logger.Info("Starting nats server...")

	for {
		var order dto.Order
		err := order.UnmarshalJSON([]byte(constants.TestData))
		if err != nil {
			n.logger.Error("Error unmarshal test data", zap.Error(err))
			return
		}

		UID, err := generateUID()
		if err != nil {
			n.logger.Error("Error generate UID", zap.Error(err))
		}

		order.OrderUID = UID
		data, _ := order.MarshalJSON()

		err = n.sc.Publish(viper.GetString("nats.subject"), data)
		if err != nil {
			n.logger.Error("Error Publish message", zap.Error(err))
			return
		}

		time.Sleep(5 * time.Second)
	}
}

func generateUID() (string, error) {
	uid := make([]byte, 16)

	_, err := rand.Read(uid)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(uid), nil
}
