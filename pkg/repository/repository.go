package repository

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/model"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

type Delivery interface {
	Create(delivery *model.Delivery) (int, error)
	GetAll() ([]model.Delivery, error)
	GetById(id int) (*model.Delivery, error)
	Update(updatedDelivery *model.Delivery) error
	Delete(id int) error
}

type Item interface {
	Create(item *model.Item) (int, error)
	GetAll() ([]model.Item, error)
	GetByOrderId(orderId int) ([]model.Item, error)
	GetById(id int) (*model.Item, error)
	Update(updatedItem *model.Item) error
	Delete(id int) error
}

type Order interface {
	Create(order *model.Order) (int, error)
	GetById(id int) (*model.Order, error)
	Update(updatedOrder *model.Order) error
	Delete(id int) error
}

type Payment interface {
	Create(payment *model.Payment) (int, error)
	GetAll() ([]model.Payment, error)
	GetById(id int) (*model.Payment, error)
	Update(updatedPayment *model.Payment) error
	Delete(id int) error
}

type Customer interface {
	Create(customer *model.Customer) error
	GetAll() ([]model.Customer, error)
	GetByUid(customerUid string) (*model.Customer, error)
	Update(updatedCustomer *model.Customer) error
	Delete(customerUid string) error
}

type Store interface {
	Set(ctx context.Context, key int, value interface{}, lifetime time.Duration) error
	Get(ctx context.Context, key int) (interface{}, error)
	Delete(ctx context.Context, key int) error
}

type Repository struct {
	Delivery
	Item
	Order
	Payment
	Customer
	Store
}

func NewRepository(db *sqlx.DB, redisClient *redis.Client, log *zap.Logger) *Repository {
	return &Repository{
		Delivery: NewDeliveryRepository(db, log),
		Item:     NewItemRepository(db, log),
		Order:    NewOrderRepository(db, log),
		Payment:  NewPaymentRepository(db, log),
		Customer: NewCustomerRepository(db, log),
		Store:    NewRedisStore(redisClient),
	}
}
