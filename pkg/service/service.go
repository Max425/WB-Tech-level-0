package service

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/Max425/WB-Tech-level-0/pkg/repository"
)

type Order interface {
	CreateOrder(ctx context.Context, order *core.Order) (int, error)
	GetCustomerOrders(ctx context.Context, customerId string) ([]core.Order, error)
	GetOrderById(ctx context.Context, id int) (*core.Order, error)
}

type Service struct {
	Order Order
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repo),
	}
}
