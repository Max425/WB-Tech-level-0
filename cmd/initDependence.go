package initial

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/repository"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strconv"
)

func InitConfig() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatalf("Ошибка при загрузке файла .env: %s", err)
	}

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func InitLogger() (*zap.Logger, error) {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	return config.Build()
}

func InitPostgres(ctx context.Context) (*sqlx.DB, error) {
	return repository.NewPostgresDB(repository.PostgresConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("PG_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
}

func InitRedis() (*redis.Client, error) {
	redisDb, err := strconv.Atoi(viper.GetString("redis.db"))
	if err != nil {
		return nil, err
	}
	return repository.NewRedisClient(repository.RedisConfig{
		Addr:     viper.GetString("redis.addr"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDb,
	})
}
