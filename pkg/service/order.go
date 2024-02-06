package service

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/constants"
	"github.com/Max425/WB-Tech-level-0/pkg/model"
	"github.com/Max425/WB-Tech-level-0/pkg/repository"
	"go.uber.org/zap"
)

type OrderService struct {
	repo repository.Repository
	log  *zap.Logger
}

func NewOrderService(repo repository.Repository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *model.Order) (int, error) {
	id, err := s.repo.Order.Create(order)
	if err != nil {
		s.log.Error("Error CreateOrder", zap.Error(err))
		return 0, err
	}
	order.ID = id
	err = s.repo.Store.Set(ctx, order.ID, order, constants.CacheDuration)
	if err != nil {
		s.log.Error("Error set to cache", zap.Error(err))
	}

	return id, nil
}

func (s *OrderService) GetOrderById(ctx context.Context, id int) (*model.Order, error) {
	cache, err := s.repo.Store.Get(ctx, id)
	orderCache := cache.(*model.Order)
	if err == nil && orderCache != nil {
		s.log.Info("Order from cache")
		return orderCache, nil
	}
	order, err := s.repo.Order.GetById(id)
	if err != nil {
		s.log.Error("Error set to cache", zap.Error(err))
		return nil, err
	}

	payment, err := s.repo.Payment.GetById(order.Payment.ID)
	if err != nil {
		s.log.Error("Error get payment for order", zap.Error(err))
		return nil, err
	}
	order.Payment = *payment

	delivery, err := s.repo.Delivery.GetById(order.Delivery.ID)
	if err != nil {
		s.log.Error("Error get delivery for order", zap.Error(err))
		return nil, err
	}
	order.Delivery = *delivery

	items, err := s.repo.Item.GetByOrderId(order.ID)
	if err != nil {
		s.log.Error("Error get items for order", zap.Error(err))
		return nil, err
	}
	order.Items = items

	return order, nil
}
