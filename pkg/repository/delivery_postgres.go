package repository

import (
	"fmt"
	"github.com/Max425/WB-Tech-level-0/pkg/constants"
	"github.com/Max425/WB-Tech-level-0/pkg/model"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type DeliveryRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewDeliveryRepository(db *sqlx.DB, log *zap.Logger) *DeliveryRepository {
	return &DeliveryRepository{db: db, log: log}
}

func (r *DeliveryRepository) Create(delivery *model.Delivery) (int, error) {
	var id int

	query := fmt.Sprintf(`
		INSERT INTO %s (name, phone, zip, city, address, region, email)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING delivery_id`,
		constants.DeliveryTable)

	err := r.db.QueryRow(query, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email).Scan(&id)
	if err != nil {
		r.log.Error("Error creating delivery", zap.Error(err))
		return 0, err
	}

	return id, nil
}

func (r *DeliveryRepository) GetAll() ([]model.Delivery, error) {
	var deliveries []model.Delivery

	query := fmt.Sprintf("SELECT * FROM %s", constants.DeliveryTable)
	err := r.db.Select(&deliveries, query)
	if err != nil {
		r.log.Error("Error retrieving deliveries", zap.Error(err))
		return nil, err
	}

	return deliveries, nil
}

func (r *DeliveryRepository) GetById(id int) (*model.Delivery, error) {
	var delivery model.Delivery

	query := fmt.Sprintf("SELECT * FROM %s WHERE delivery_id = $1", constants.DeliveryTable)
	err := r.db.Get(&delivery, query, id)
	if err != nil {
		r.log.Error("Error retrieving delivery by ID", zap.Error(err))
		return nil, err
	}

	return &delivery, nil
}

func (r *DeliveryRepository) Update(id int, updatedDelivery *model.Delivery) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET name=$1, phone=$2, zip=$3, city=$4, address=$5, region=$6, email=$7
		WHERE delivery_id = $8`,
		constants.DeliveryTable)

	_, err := r.db.Exec(query, updatedDelivery.Name, updatedDelivery.Phone, updatedDelivery.Zip,
		updatedDelivery.City, updatedDelivery.Address, updatedDelivery.Region, updatedDelivery.Email, id)
	if err != nil {
		r.log.Error("Error updating delivery", zap.Error(err))
		return err
	}

	return nil
}

func (r *DeliveryRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE delivery_id = $1", constants.DeliveryTable)

	_, err := r.db.Exec(query, id)
	if err != nil {
		r.log.Error("Error deleting delivery", zap.Error(err))
		return err
	}

	return nil
}
