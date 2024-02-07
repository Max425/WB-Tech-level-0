package service

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/repository"
	"go.uber.org/zap"
)

type Order interface {
	CreateOrder(ctx context.Context, data []byte) (int, error)
	GetOrderByUID(ctx context.Context, UID string) (string, error)
	GetCustomerOrders(customerUID string) ([]string, error)
	LoadOrdersToCache(ctx context.Context) error
}

type Service struct {
	Order Order
}

func NewService(repo *repository.Repository, log *zap.Logger) *Service {
	return &Service{
		Order: NewOrderService(repo.Order, repo.Store, log),
	}
}
