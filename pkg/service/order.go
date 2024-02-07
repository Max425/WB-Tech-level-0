package service

import (
	"context"
	"fmt"
	"github.com/Max425/WB-Tech-level-0/pkg/constants"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/Max425/WB-Tech-level-0/pkg/model/dto"
	"github.com/Max425/WB-Tech-level-0/pkg/repository"
	"go.uber.org/zap"
)

type OrderService struct {
	repoOrder repository.Order
	repoStore repository.Store
	log       *zap.Logger
}

func NewOrderService(repoOrder repository.Order, repoStore repository.Store, log *zap.Logger) *OrderService {
	return &OrderService{repoOrder: repoOrder, repoStore: repoStore, log: log}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *dto.Order, data []byte) (int, error) {
	id, err := s.repoOrder.Create(&core.Order{OrderUID: order.OrderUID, Data: data})
	if err != nil {
		s.log.Error(fmt.Sprintf("Error create order with UID: %s", order.OrderUID), zap.Error(err))
		return 0, err
	}

	err = s.repoStore.Set(ctx, order.OrderUID, data, constants.CacheDuration)
	if err != nil {
		s.log.Error("Error set order to cache", zap.Error(err))
	}

	return id, nil
}

func (s *OrderService) GetOrderByUID(ctx context.Context, UID string) (string, error) {
	cache, err := s.repoStore.Get(ctx, UID)
	if err == nil {
		s.log.Info("Order from cache")
		return string(cache), nil
	}
	order, err := s.repoOrder.GetByUID(UID)
	if err != nil {
		s.log.Error("Error set to cache", zap.Error(err))
		return "", err
	}

	return string(order.Data), nil
}

func (s *OrderService) GetCustomerOrders(customerUID string) ([]string, error) {
	orders, err := s.repoOrder.GetCustomerOrders(customerUID)
	if err != nil {
		s.log.Error("Error get customer orders", zap.Error(err))
		return nil, err
	}
	var customerOrders []string
	for _, order := range orders {
		customerOrders = append(customerOrders, string(order.Data))
	}

	return customerOrders, nil
}

func (s *OrderService) LoadOrdersToCache(ctx context.Context) error {
	orders, err := s.repoOrder.GetAll()
	if err != nil {
		s.log.Error("Error get all orders", zap.Error(err))
		return err
	}
	for _, order := range orders {
		err = s.repoStore.Set(ctx, order.OrderUID, order.Data, constants.CacheDuration)
		if err != nil {
			s.log.Error("Error load order to cache", zap.Error(err))
		}
	}
	return nil
}
