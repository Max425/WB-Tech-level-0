package repository

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

type Order interface {
	Create(order *core.Order) (int, error)
	GetByUID(UID string) (*core.Order, error)
	GetCustomerOrders(customerUID string) ([]core.Order, error)
	GetAll() ([]core.Order, error)
	DeleteByUID(UID string) error
}

type Store interface {
	Set(ctx context.Context, key string, value []byte, lifetime time.Duration) error
	Get(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
}

type Repository struct {
	Order
	Store
}

func NewRepository(db *sqlx.DB, redisClient *redis.Client, log *zap.Logger) *Repository {
	return &Repository{
		Order: NewOrderRepository(db, log),
		Store: NewRedisStore(redisClient),
	}
}
