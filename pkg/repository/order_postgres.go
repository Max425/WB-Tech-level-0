package repository

import (
	"database/sql"
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

	tx, err := r.db.Begin()
	if err != nil {
		r.log.Error("Error beginning transaction", zap.Error(err))
		return 0, err
	}

	query := fmt.Sprintf(`
		INSERT INTO %s (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature,
			customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`,
		constants.OrderTable)

	err = tx.QueryRow(query, order.OrderUID, order.TrackNumber, order.Entry, order.Delivery.ID,
		order.Payment.ID, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService,
		order.ShardKey, order.SMID, order.DateCreated, order.OofShard).Scan(&id)
	if err != nil || addOrderItems(tx, order.ID, order.Items) != nil {
		r.log.Error("Error creating order", zap.Error(err))
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *OrderRepository) GetById(id int) (*core.Order, error) {
	var order core.Order

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", constants.OrderTable)
	err := r.db.Get(&order, query, id)
	if err != nil {
		r.log.Error("Error retrieving order by ID", zap.Error(err))
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) GetCustomerOrders(customerId string) ([]core.Order, error) {
	var orders []core.Order

	query := fmt.Sprintf("SELECT * FROM %s WHERE customer_id = $1", constants.OrderTable)
	err := r.db.Get(&orders, query, customerId)
	if err != nil {
		r.log.Error("Error retrieving order by Customer ID", zap.Error(err))
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) Update(updatedOrder *core.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		r.log.Error("Error beginning transaction", zap.Error(err))
		return err
	}

	query := fmt.Sprintf(`
		UPDATE %s
		SET order_uid=$1, track_number=$2, entry=$3, delivery_id=$4, payment_id=$5, locale=$6,
		internal_signature=$7, customer_id=$8, delivery_service=$9, shard_key=$10, sm_id=$11,
		date_created=$12, oof_shard=$13 WHERE id = $14`,
		constants.OrderTable)

	_, err = tx.Exec(query, updatedOrder.OrderUID, updatedOrder.TrackNumber, updatedOrder.Entry,
		updatedOrder.Delivery.ID, updatedOrder.Payment.ID, updatedOrder.Locale, updatedOrder.InternalSignature,
		updatedOrder.CustomerID, updatedOrder.DeliveryService, updatedOrder.ShardKey, updatedOrder.SMID,
		updatedOrder.DateCreated, updatedOrder.OofShard, updatedOrder.ID)

	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE order_id = $1", constants.OrderItemTable)
	_, err = tx.Exec(deleteQuery, updatedOrder.ID)
	if err != nil {
		tx.Rollback()
		r.log.Error("Error deleting order items", zap.Error(err))
		return err
	}
	if err != nil || addOrderItems(tx, updatedOrder.ID, updatedOrder.Items) != nil {
		tx.Rollback()
		r.log.Error("Error updating order", zap.Error(err))
		return err
	}

	return tx.Commit()
}

func (r *OrderRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", constants.OrderTable)

	_, err := r.db.Exec(query, id)
	if err != nil {
		r.log.Error("Error deleting order", zap.Error(err))
		return err
	}

	return nil
}

func addOrderItems(tx *sql.Tx, orderID int, updatedItems []core.Item) error {
	var query strings.Builder
	query.WriteString(fmt.Sprintf("INSERT INTO %s (order_id, item_id) VALUES ", constants.OrderItemTable))

	var values []interface{}

	for i, item := range updatedItems {
		if i > 0 {
			query.WriteString(", ")
		}
		query.WriteString(fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		values = append(values, orderID, item.ID)
	}

	insertQuery := query.String()
	_, err := tx.Exec(insertQuery, values...)
	if err != nil {
		return err
	}

	return nil
}
