package repository

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

type Delivery interface {
	Create(delivery *core.Delivery) (int, error)
	GetAll() ([]core.Delivery, error)
	GetById(id int) (*core.Delivery, error)
	Update(updatedDelivery *core.Delivery) error
	Delete(id int) error
}

type Item interface {
	Create(item *core.Item) (int, error)
	GetAll() ([]core.Item, error)
	GetByOrderId(orderId int) ([]core.Item, error)
	GetById(id int) (*core.Item, error)
	Update(updatedItem *core.Item) error
	Delete(id int) error
}

type Order interface {
	Create(order *core.Order) (int, error)
	GetById(id int) (*core.Order, error)
	GetCustomerOrders(customerId string) ([]core.Order, error)
	GetAll() ([]core.Order, error)
	Update(updatedOrder *core.Order) error
	Delete(id int) error
}

type Payment interface {
	Create(payment *core.Payment) (int, error)
	GetAll() ([]core.Payment, error)
	GetById(id int) (*core.Payment, error)
	Update(updatedPayment *core.Payment) error
	Delete(id int) error
}

type Customer interface {
	Create(customer *core.Customer) error
	GetAll() ([]core.Customer, error)
	GetByUid(customerUid string) (*core.Customer, error)
	Update(updatedCustomer *core.Customer) error
	Delete(customerUid string) error
}

type Store interface {
	Set(ctx context.Context, key int, value []byte, lifetime time.Duration) error
	Get(ctx context.Context, key int) ([]byte, error)
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
