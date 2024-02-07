package main

import (
	"crypto/rand"
	"encoding/hex"
	initial "github.com/Max425/WB-Tech-level-0/cmd"
	"github.com/Max425/WB-Tech-level-0/pkg/constants"
	"github.com/Max425/WB-Tech-level-0/pkg/model/dto"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"time"
)

func main() {
	initial.InitConfig()
	logger, err := initial.InitLogger()
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Starting nats server...")

	sc, err := stan.Connect(
		viper.GetString("nats.clusterID"),
		viper.GetString("nats.serverID"),
		stan.NatsURL(viper.GetString("nats.url")),
	)
	if err != nil {
		logger.Error("Error stan Connect", zap.Error(err))
		return
	}

	for {
		var order dto.Order
		err = order.UnmarshalJSON([]byte(constants.TestData))
		if err != nil {
			logger.Error("Error unmarshal test data", zap.Error(err))
			return
		}

		UID, err := generateUID()
		if err != nil {
			logger.Error("Error generate UID", zap.Error(err))
		}

		order.OrderUID = UID
		data, _ := order.MarshalJSON()

		err = sc.Publish(viper.GetString("nats.subject"), data)
		if err != nil {
			logger.Error("Error Publish message", zap.Error(err))
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
