package repository

import (
	"fmt"
	"github.com/Max425/WB-Tech-level-0/pkg/constants"
	"github.com/Max425/WB-Tech-level-0/pkg/model"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type ItemRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewItemRepository(db *sqlx.DB, log *zap.Logger) *ItemRepository {
	return &ItemRepository{db: db, log: log}
}

func (r *ItemRepository) Create(item *model.Item) (int, error) {
	var id int

	query := fmt.Sprintf(`
		INSERT INTO %s (chrt_id, track_number, price, rid, item_name, sale, size, total_price, nm_id, brand, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`,
		constants.ItemTable)

	err := r.db.QueryRow(query, item.ChrtID, item.TrackNumber, item.Price, item.RID, item.Name, item.Sale,
		item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status).Scan(&id)
	if err != nil {
		r.log.Error("Error creating item", zap.Error(err))
		return 0, err
	}

	return id, nil
}

func (r *ItemRepository) GetAll() ([]model.Item, error) {
	var items []model.Item

	query := fmt.Sprintf("SELECT * FROM %s", constants.ItemTable)
	err := r.db.Select(&items, query)
	if err != nil {
		r.log.Error("Error retrieving items", zap.Error(err))
		return nil, err
	}

	return items, nil
}

func (r *ItemRepository) GetByOrderId(orderId int) ([]model.Item, error) {
	var items []model.Item

	query := fmt.Sprintf("SELECT * FROM %s WHERE id IN (SELECT item_id FROM %s WHERE order_id = $1)", constants.ItemTable, constants.OrderItemTable)
	err := r.db.Select(&items, query, orderId)
	if err != nil {
		r.log.Error("Error retrieving items", zap.Error(err))
		return nil, err
	}

	return items, nil
}

func (r *ItemRepository) GetById(id int) (*model.Item, error) {
	var item model.Item

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", constants.ItemTable)
	err := r.db.Get(&item, query, id)
	if err != nil {
		r.log.Error("Error retrieving item by ID", zap.Error(err))
		return nil, err
	}

	return &item, nil
}

func (r *ItemRepository) Update(updatedItem *model.Item) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET chrt_id=$1, track_number=$2, price=$3, rid=$4, item_name=$5, sale=$6, size=$7, total_price=$8,
		nm_id=$9, brand=$10, status=$11 WHERE id = $12`,
		constants.ItemTable)

	_, err := r.db.Exec(query, updatedItem.ChrtID, updatedItem.TrackNumber, updatedItem.Price, updatedItem.RID,
		updatedItem.Name, updatedItem.Sale, updatedItem.Size, updatedItem.TotalPrice, updatedItem.NmID,
		updatedItem.Brand, updatedItem.Status, updatedItem.ID)
	if err != nil {
		r.log.Error("Error updating item", zap.Error(err))
		return err
	}

	return nil
}

func (r *ItemRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", constants.ItemTable)

	_, err := r.db.Exec(query, id)
	if err != nil {
		r.log.Error("Error deleting item", zap.Error(err))
		return err
	}

	return nil
}
