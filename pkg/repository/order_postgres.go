package repository

import (
	"database/sql"
	"fmt"
	"github.com/Max425/WB-Tech-level-0/pkg/constants"
	"github.com/Max425/WB-Tech-level-0/pkg/model"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type OrderRepository struct {
	db           *sqlx.DB
	log          *zap.Logger
	deliveryRepo *DeliveryRepository
	paymentRepo  *PaymentRepository
	itemRepo     *ItemRepository
}

func NewOrderRepository(db *sqlx.DB, log *zap.Logger, deliveryRepo *DeliveryRepository, paymentRepo *PaymentRepository, itemRepo *ItemRepository) *OrderRepository {
	return &OrderRepository{db: db, log: log, deliveryRepo: deliveryRepo, paymentRepo: paymentRepo, itemRepo: itemRepo}
}

func (r *OrderRepository) Create(order *model.Order) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		r.log.Error("Error beginning transaction", zap.Error(err))
		return 0, err
	}

	// Create Delivery
	deliveryID, err := r.deliveryRepo.Create(&order.Delivery)
	if err != nil {
		tx.Rollback()
		r.log.Error("Error creating delivery", zap.Error(err))
		return 0, err
	}

	// Create Payment
	paymentID, err := r.paymentRepo.Create(&order.Payment)
	if err != nil {
		tx.Rollback()
		r.log.Error("Error creating payment", zap.Error(err))
		return 0, err
	}

	// Create Order
	query := fmt.Sprintf(`
		INSERT INTO %s (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature,
		customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`,
		constants.OrderTable)

	err = tx.QueryRow(query, order.OrderUID, order.TrackNumber, order.Entry, deliveryID, paymentID, order.Locale,
		order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SMID, order.DateCreated, order.OofShard).Scan(&order.ID)
	if err != nil {
		tx.Rollback()
		r.log.Error("Error creating order", zap.Error(err))
		return 0, err
	}

	// Create Order Items
	for _, item := range order.Items {
		itemID, err := r.itemRepo.Create(&item)
		if err != nil {
			tx.Rollback()
			r.log.Error("Error creating order item", zap.Error(err))
			return 0, err
		}

		// Link Order and Item
		linkOrderItemQuery := fmt.Sprintf("INSERT INTO %s (order_id, item_id) VALUES ($1, $2)", constants.OrderItemTable)
		_, err = tx.Exec(linkOrderItemQuery, order.ID, itemID)
		if err != nil {
			tx.Rollback()
			r.log.Error("Error linking order and item", zap.Error(err))
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		r.log.Error("Error committing transaction", zap.Error(err))
		return 0, err
	}

	return order.ID, nil
}

func (r *OrderRepository) GetAll() ([]model.Order, error) {
	var orders []model.Order

	query := fmt.Sprintf("SELECT * FROM %s", constants.OrderTable)
	err := r.db.Select(&orders, query)
	if err != nil {
		r.log.Error("Error retrieving orders", zap.Error(err))
		return nil, err
	}

	for i := range orders {
		// Fetch and set nested entities
		delivery, err := r.deliveryRepo.GetById(orders[i].Delivery.ID)
		if err != nil {
			r.log.Error("Error retrieving delivery for order", zap.Error(err))
			return nil, err
		}
		orders[i].Delivery = *delivery

		payment, err := r.paymentRepo.GetById(orders[i].Payment.ID)
		if err != nil {
			r.log.Error("Error retrieving payment for order", zap.Error(err))
			return nil, err
		}
		orders[i].Payment = *payment

		items, err := r.getItemsByOrderId(orders[i].ID)
		if err != nil {
			r.log.Error("Error retrieving items for order", zap.Error(err))
			return nil, err
		}
		orders[i].Items = items
	}

	return orders, nil
}

func (r *OrderRepository) GetById(id int) (*model.Order, error) {
	var order model.Order

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", constants.OrderTable)
	err := r.db.Get(&order, query, id)
	if err != nil {
		r.log.Error("Error retrieving order by ID", zap.Error(err))
		return nil, err
	}

	// Fetch and set nested entities
	delivery, err := r.deliveryRepo.GetById(order.Delivery.ID)
	if err != nil {
		r.log.Error("Error retrieving delivery for order", zap.Error(err))
		return nil, err
	}
	order.Delivery = *delivery

	payment, err := r.paymentRepo.GetById(order.Payment.ID)
	if err != nil {
		r.log.Error("Error retrieving payment for order", zap.Error(err))
		return nil, err
	}
	order.Payment = *payment

	items, err := r.getItemsByOrderId(order.ID)
	if err != nil {
		r.log.Error("Error retrieving items for order", zap.Error(err))
		return nil, err
	}
	order.Items = items

	return &order, nil
}

func (r *OrderRepository) Update(id int, updatedOrder *model.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		r.log.Error("Error beginning transaction", zap.Error(err))
		return err
	}

	// Update Delivery
	err = r.deliveryRepo.Update(updatedOrder.Delivery.ID, &updatedOrder.Delivery)
	if err != nil {
		tx.Rollback()
		r.log.Error("Error updating delivery", zap.Error(err))
		return err
	}

	// Update Payment
	err = r.paymentRepo.Update(updatedOrder.Payment.ID, &updatedOrder.Payment)
	if err != nil {
		tx.Rollback()
		r.log.Error("Error updating payment", zap.Error(err))
		return err
	}

	// Update Order
	query := fmt.Sprintf(`
		UPDATE %s
		SET order_uid=$1, track_number=$2, entry=$3, locale=$4, internal_signature=$5, customer_id=$6,
		delivery_service=$7, shard_key=$8, sm_id=$9, date_created=$10, oof_shard=$11
		WHERE id = $12`,
		constants.OrderTable)

	_, err = tx.Exec(query, updatedOrder.OrderUID, updatedOrder.TrackNumber, updatedOrder.Entry,
		updatedOrder.Locale, updatedOrder.InternalSignature, updatedOrder.CustomerID,
		updatedOrder.DeliveryService, updatedOrder.ShardKey, updatedOrder.SMID,
		updatedOrder.DateCreated, updatedOrder.OofShard, id)
	if err != nil {
		tx.Rollback()
		r.log.Error("Error updating order", zap.Error(err))
		return err
	}

	// Update Order Items
	err = r.updateOrderItems(tx, id, updatedOrder.Items)
	if err != nil {
		tx.Rollback()
		r.log.Error("Error updating order items", zap.Error(err))
		return err
	}

	err = tx.Commit()
	if err != nil {
		r.log.Error("Error committing transaction", zap.Error(err))
		return err
	}

	return nil
}

func (r *OrderRepository) Delete(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		r.log.Error("Error beginning transaction", zap.Error(err))
		return err
	}

	// Delete Order Items
	deleteOrderItemsQuery := fmt.Sprintf("DELETE FROM %s WHERE order_id = $1", constants.OrderItemTable)
	_, err = tx.Exec(deleteOrderItemsQuery, id)
	if err != nil {
		tx.Rollback()
		r.log.Error("Error deleting order items", zap.Error(err))
		return err
	}

	// Delete Order
	deleteOrderQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", constants.OrderTable)
	_, err = tx.Exec(deleteOrderQuery, id)
	if err != nil {
		tx.Rollback()
		r.log.Error("Error deleting order", zap.Error(err))
		return err
	}

	err = tx.Commit()
	if err != nil {
		r.log.Error("Error committing transaction", zap.Error(err))
		return err
	}

	return nil
}

// Helper method to get items for a specific order
func (r *OrderRepository) getItemsByOrderId(orderID int) ([]model.Item, error) {
	var items []model.Item

	query := fmt.Sprintf(`
		SELECT i.*
		FROM %s oi
		JOIN %s i ON oi.item_id = i.id
		WHERE oi.order_id = $1`,
		constants.OrderItemTable, constants.ItemTable)

	err := r.db.Select(&items, query, orderID)
	if err != nil {
		r.log.Error("Error retrieving items for order", zap.Error(err))
		return nil, err
	}

	return items, nil
}

// Helper method to update order items
func (r *OrderRepository) updateOrderItems(tx *sql.Tx, orderID int, updatedItems []model.Item) error {
	// Delete existing order items
	deleteOrderItemsQuery := fmt.Sprintf("DELETE FROM %s WHERE order_id = $1", constants.OrderItemTable)
	_, err := tx.Exec(deleteOrderItemsQuery, orderID)
	if err != nil {
		return err
	}

	// Create new order items
	for _, item := range updatedItems {
		itemID, err := r.itemRepo.Create(&item)
		if err != nil {
			return err
		}

		// Link Order and Item
		linkOrderItemQuery := fmt.Sprintf("INSERT INTO %s (order_id, item_id) VALUES ($1, $2)", constants.OrderItemTable)
		_, err = tx.Exec(linkOrderItemQuery, orderID, itemID)
		if err != nil {
			return err
		}
	}

	return nil
}
