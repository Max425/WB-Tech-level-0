package service

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/model"
	"github.com/Max425/WB-Tech-level-0/pkg/repository"
)

type Order interface {
	CreateOrder(ctx context.Context, order *model.Order) (int, error)
	GetOrderById(ctx context.Context, id int) (*model.Order, error)
}

type Service struct {
	Order Order
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repo),
	}
}
