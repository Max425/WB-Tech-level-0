package service

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/constants"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/Max425/WB-Tech-level-0/pkg/repository"
	"go.uber.org/zap"
)

type OrderService struct {
	repo *repository.Repository
	log  *zap.Logger
}

func NewOrderService(repo *repository.Repository, log *zap.Logger) *OrderService {
	return &OrderService{repo: repo, log: log}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *core.Order) (int, error) {
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

func (s *OrderService) GetOrderById(ctx context.Context, id int) (*core.Order, error) {
	cache, err := s.repo.Store.Get(ctx, id)
	orderCache := cache.(*core.Order)
	if err == nil && orderCache != nil {
		s.log.Info("Order from cache")
		return orderCache, nil
	}
	order, err := s.repo.Order.GetById(id)
	if err != nil {
		s.log.Error("Error set to cache", zap.Error(err))
		return nil, err
	}

	return s.fillOrder(order)
}

func (s *OrderService) GetCustomerOrders(ctx context.Context, customerId string) ([]core.Order, error) {
	orders, err := s.repo.Order.GetCustomerOrders(customerId)
	if err != nil {
		s.log.Error("Error get customer orders", zap.Error(err))
		return nil, err
	}
	for i := 0; i < len(orders); i++ {
		cache, err := s.repo.Store.Get(ctx, orders[i].ID)
		orderCache := cache.(*core.Order)
		if err == nil && orderCache != nil {
			orders[i] = *orderCache
		} else {
			data, err := s.fillOrder(&orders[i])
			if err != nil {
				return nil, err
			}
			orders[i] = *data
		}
	}

	return orders, nil
}

func (s *OrderService) LoadOrdersToCache(ctx context.Context) error {
	orders, err := s.repo.Order.GetAll()
	if err != nil {
		s.log.Error("Error load orders to cache", zap.Error(err))
		return err
	}
	for _, order := range orders {
		data, err := s.fillOrder(&order)
		if err != nil {
			s.log.Error("Error load orders to cache", zap.Error(err))
			return err
		}
		err = s.repo.Store.Set(ctx, data.ID, data, constants.CacheDuration)
		if err != nil {
			s.log.Error("Error load order to cache", zap.Error(err))
		}
	}
	return nil
}

func (s *OrderService) fillOrder(order *core.Order) (*core.Order, error) {
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
