package repository

import (
	"fmt"
	"github.com/Max425/WB-Tech-level-0/pkg/constants"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strings"
)

type OrderRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewOrderRepository(db *sqlx.DB, log *zap.Logger) *OrderRepository {
	return &OrderRepository{db: db, log: log}
}

func (r *OrderRepository) Create(order *core.Order) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (order_uid, data) VALUES ($1, $2) RETURNING id`,
		constants.OrderTable)

	err := r.db.QueryRow(query, order.OrderUID, order.Data).Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			r.log.Error(fmt.Sprintf("Order with UID %s already created", order.OrderUID), zap.Error(err))
			return 0, constants.AlreadyExistsError
		}
		r.log.Error("Error create order", zap.Error(err))
		return 0, err
	}

	return id, nil
}

func (r *OrderRepository) GetByUID(UID string) (*core.Order, error) {
	var order *core.Order

	query := fmt.Sprintf("SELECT id, order_uid, data FROM %s WHERE order_uid = $1",
		constants.OrderTable)

	err := r.db.Get(order, query, UID)
	if err != nil {
		r.log.Error("Error get order by UID", zap.Error(err))
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) GetCustomerOrders(customerUID string) ([]core.Order, error) {
	var orders []core.Order

	query := fmt.Sprintf("SELECT id, order_uid, data FROM %s WHERE data->>'customer_id' = $1", constants.OrderTable)
	err := r.db.Select(&orders, query, customerUID)
	if err != nil {
		r.log.Error(fmt.Sprintf("Error get orders by Customer ID = %s", customerUID), zap.Error(err))
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetAll() ([]core.Order, error) {
	var orders []core.Order

	query := fmt.Sprintf("SELECT id, order_uid, data FROM %s", constants.OrderTable)
	err := r.db.Get(&orders, query)
	if err != nil {
		r.log.Error("Error get all orders", zap.Error(err))
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) DeleteByUID(UID string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE order_uid = $1", constants.OrderTable)

	_, err := r.db.Exec(query, UID)
	if err != nil {
		r.log.Error(fmt.Sprintf("Error deleting order with UID %s", UID), zap.Error(err))
		return err
	}

	return nil
}
