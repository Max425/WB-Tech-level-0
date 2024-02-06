package service

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/Max425/WB-Tech-level-0/pkg/repository"
	"go.uber.org/zap"
)

type PaymentService struct {
	repo *repository.Repository
	log  *zap.Logger
}

func NewPaymentService(repo *repository.Repository, log *zap.Logger) *PaymentService {
	return &PaymentService{repo: repo, log: log}
}

func (s *PaymentService) Create(ctx context.Context, payment *core.Payment) (int, error) {
	id, err := s.repo.Payment.Create(payment)
	if err != nil {
		s.log.Error("Error creating payment", zap.Error(err))
		return 0, err
	}
	return id, nil
}

func (s *PaymentService) GetAll() ([]core.Payment, error) {
	return s.repo.Payment.GetAll()
}

func (s *PaymentService) GetById(ctx context.Context, id int) (*core.Payment, error) {
	return s.repo.Payment.GetById(id)
}

func (s *PaymentService) Update(ctx context.Context, updatedPayment *core.Payment) error {
	return s.repo.Payment.Update(updatedPayment)
}

func (s *PaymentService) Delete(id int) error {
	return s.repo.Payment.Delete(id)
}
