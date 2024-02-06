package repository

import (
	"fmt"
	"github.com/Max425/WB-Tech-level-0/pkg/constants"
	"github.com/Max425/WB-Tech-level-0/pkg/model"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type PaymentRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewPaymentRepository(db *sqlx.DB, log *zap.Logger) *PaymentRepository {
	return &PaymentRepository{db: db, log: log}
}

func (r *PaymentRepository) Create(payment *model.Payment) (int, error) {
	var id int

	query := fmt.Sprintf(`
		INSERT INTO %s (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`,
		constants.PaymentTable)

	err := r.db.QueryRow(query, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount,
		payment.PaymentDT, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee).Scan(&id)
	if err != nil {
		r.log.Error("Error creating payment", zap.Error(err))
		return 0, err
	}

	return id, nil
}

func (r *PaymentRepository) GetAll() ([]model.Payment, error) {
	var payments []model.Payment

	query := fmt.Sprintf("SELECT * FROM %s", constants.PaymentTable)
	err := r.db.Select(&payments, query)
	if err != nil {
		r.log.Error("Error retrieving payments", zap.Error(err))
		return nil, err
	}

	return payments, nil
}

func (r *PaymentRepository) GetById(id int) (*model.Payment, error) {
	var payment model.Payment

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", constants.PaymentTable)
	err := r.db.Get(&payment, query, id)
	if err != nil {
		r.log.Error("Error retrieving payment by ID", zap.Error(err))
		return nil, err
	}

	return &payment, nil
}

func (r *PaymentRepository) Update(updatedPayment *model.Payment) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET transaction=$1, request_id=$2, currency=$3, provider=$4, amount=$5, payment_dt=$6, bank=$7, 
		delivery_cost=$8, goods_total=$9, custom_fee=$10
		WHERE id = $11`,
		constants.PaymentTable)

	_, err := r.db.Exec(query, updatedPayment.Transaction, updatedPayment.RequestID, updatedPayment.Currency,
		updatedPayment.Provider, updatedPayment.Amount, updatedPayment.PaymentDT, updatedPayment.Bank,
		updatedPayment.DeliveryCost, updatedPayment.GoodsTotal, updatedPayment.CustomFee, updatedPayment.ID)
	if err != nil {
		r.log.Error("Error updating payment", zap.Error(err))
		return err
	}

	return nil
}

func (r *PaymentRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", constants.PaymentTable)

	_, err := r.db.Exec(query, id)
	if err != nil {
		r.log.Error("Error deleting payment", zap.Error(err))
		return err
	}

	return nil
}
