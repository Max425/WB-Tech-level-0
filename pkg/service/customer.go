package service

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/Max425/WB-Tech-level-0/pkg/repository"
	"go.uber.org/zap"
)

type CustomerService struct {
	repo *repository.Repository
	log  *zap.Logger
}

func NewCustomerService(repo *repository.Repository, log *zap.Logger) *CustomerService {
	return &CustomerService{repo: repo, log: log}
}

func (s *CustomerService) Create(ctx context.Context, customer *core.Customer) error {
	err := s.repo.Customer.Create(customer)
	if err != nil {
		s.log.Error("Error creating customer", zap.Error(err))
		return err
	}
	return nil
}

func (s *CustomerService) GetAll() ([]core.Customer, error) {
	return s.repo.Customer.GetAll()
}

func (s *CustomerService) GetByUid(ctx context.Context, customerUid string) (*core.Customer, error) {
	return s.repo.Customer.GetByUid(customerUid)
}

func (s *CustomerService) Update(ctx context.Context, updatedCustomer *core.Customer) error {
	return s.repo.Customer.Update(updatedCustomer)
}

func (s *CustomerService) Delete(customerUid string) error {
	return s.repo.Customer.Delete(customerUid)
}
