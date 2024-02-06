package service

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/constants"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/Max425/WB-Tech-level-0/pkg/repository"
	"go.uber.org/zap"
)

type ItemService struct {
	repo *repository.Repository
	log  *zap.Logger
}

func NewItemService(repo *repository.Repository, log *zap.Logger) *ItemService {
	return &ItemService{repo: repo, log: log}
}

func (s *ItemService) CreateItem(ctx context.Context, item *core.Item) (int, error) {
	id, err := s.repo.Item.Create(item)
	if err != nil {
		s.log.Error("Error CreateItem", zap.Error(err))
		return 0, err
	}
	item.ID = id
	err = s.repo.Store.Set(ctx, item.ID, item, constants.CacheDuration)
	if err != nil {
		s.log.Error("Error set to cache", zap.Error(err))
	}

	return id, nil
}

func (s *ItemService) GetAllItems() ([]core.Item, error) {
	return s.repo.Item.GetAll()
}

func (s *ItemService) GetItemsByOrderId(orderId int) ([]core.Item, error) {
	return s.repo.Item.GetByOrderId(orderId)
}

func (s *ItemService) GetItemById(ctx context.Context, id int) (*core.Item, error) {
	return s.repo.Item.GetById(id)
}

func (s *ItemService) UpdateItem(ctx context.Context, updatedItem *core.Item) error {
	return s.repo.Item.Update(updatedItem)
}

func (s *ItemService) DeleteItem(id int) error {
	return s.repo.Item.Delete(id)
}
