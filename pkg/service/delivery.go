package service

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/Max425/WB-Tech-level-0/pkg/repository"
	"go.uber.org/zap"
)

type DeliveryService struct {
	repo *repository.Repository
	log  *zap.Logger
}

func NewDeliveryService(repo *repository.Repository, log *zap.Logger) *DeliveryService {
	return &DeliveryService{repo: repo, log: log}
}

func (s *DeliveryService) Create(ctx context.Context, delivery *core.Delivery) (int, error) {
	id, err := s.repo.Delivery.Create(delivery)
	if err != nil {
		s.log.Error("Error creating delivery", zap.Error(err))
		return 0, err
	}
	return id, nil
}

func (s *DeliveryService) GetAll() ([]core.Delivery, error) {
	return s.repo.Delivery.GetAll()
}

func (s *DeliveryService) GetById(ctx context.Context, id int) (*core.Delivery, error) {
	return s.repo.Delivery.GetById(id)
}

func (s *DeliveryService) Update(ctx context.Context, updatedDelivery *core.Delivery) error {
	return s.repo.Delivery.Update(updatedDelivery)
}

func (s *DeliveryService) Delete(id int) error {
	return s.repo.Delivery.Delete(id)
}
