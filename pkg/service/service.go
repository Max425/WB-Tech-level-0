package service

import (
	"context"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/Max425/WB-Tech-level-0/pkg/repository"
	"go.uber.org/zap"
)

type Order interface {
	CreateOrder(ctx context.Context, order *core.Order) (int, error)
	GetCustomerOrders(ctx context.Context, customerId string) ([]core.Order, error)
	GetOrderById(ctx context.Context, id int) (*core.Order, error)
}

type Item interface {
	CreateItem(ctx context.Context, item *core.Item) (int, error)
	GetAllItems() ([]core.Item, error)
	GetItemsByOrderId(orderId int) ([]core.Item, error)
	GetItemById(ctx context.Context, id int) (*core.Item, error)
	UpdateItem(ctx context.Context, updatedItem *core.Item) error
	DeleteItem(id int) error
}

type Delivery interface {
	Create(ctx context.Context, delivery *core.Delivery) (int, error)
	GetAll() ([]core.Delivery, error)
	GetById(ctx context.Context, id int) (*core.Delivery, error)
	Update(ctx context.Context, updatedDelivery *core.Delivery) error
	Delete(id int) error
}

type Payment interface {
	Create(ctx context.Context, payment *core.Payment) (int, error)
	GetAll() ([]core.Payment, error)
	GetById(ctx context.Context, id int) (*core.Payment, error)
	Update(ctx context.Context, updatedPayment *core.Payment) error
	Delete(id int) error
}

type Customer interface {
	Create(ctx context.Context, customer *core.Customer) error
	GetAll() ([]core.Customer, error)
	GetByUid(ctx context.Context, customerUid string) (*core.Customer, error)
	Update(ctx context.Context, updatedCustomer *core.Customer) error
	Delete(customerUid string) error
}

type Service struct {
	Order    Order
	Item     Item
	Delivery Delivery
	Payment  Payment
	Customer Customer
}

func NewService(repo *repository.Repository, log *zap.Logger) *Service {
	return &Service{
		Order:    NewOrderService(repo, log),
		Item:     NewItemService(repo, log),
		Delivery: NewDeliveryService(repo, log),
		Payment:  NewPaymentService(repo, log),
		Customer: NewCustomerService(repo, log),
	}
}
