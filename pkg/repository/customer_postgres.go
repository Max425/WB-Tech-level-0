package repository

import (
	"fmt"
	"github.com/Max425/WB-Tech-level-0/pkg/constants"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type CustomerRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewCustomerRepository(db *sqlx.DB, log *zap.Logger) *CustomerRepository {
	return &CustomerRepository{db: db, log: log}
}

func (r *CustomerRepository) Create(customer *core.Customer) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (customer_uid, email)
		VALUES ($1, $2)`,
		constants.CustomerTable)

	_, err := r.db.Exec(query, customer.CustomerUid, customer.Email)
	if err != nil {
		r.log.Error("Error creating customer", zap.Error(err))
		return err
	}

	return nil
}

func (r *CustomerRepository) GetAll() ([]core.Customer, error) {
	var customers []core.Customer

	query := fmt.Sprintf("SELECT * FROM %s", constants.CustomerTable)
	err := r.db.Select(&customers, query)
	if err != nil {
		r.log.Error("Error retrieving customers", zap.Error(err))
		return nil, err
	}

	return customers, nil
}

func (r *CustomerRepository) GetByUid(customerUid string) (*core.Customer, error) {
	var customer core.Customer

	query := fmt.Sprintf("SELECT * FROM %s WHERE customer_uid = $1", constants.CustomerTable)
	err := r.db.Get(&customer, query, customerUid)
	if err != nil {
		r.log.Error("Error retrieving customer by UID", zap.Error(err))
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerRepository) Update(updatedCustomer *core.Customer) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET email=$1
		WHERE customer_uid = $2`,
		constants.CustomerTable)

	_, err := r.db.Exec(query, updatedCustomer.Email, updatedCustomer.CustomerUid)
	if err != nil {
		r.log.Error("Error updating customer", zap.Error(err))
		return err
	}

	return nil
}

func (r *CustomerRepository) Delete(customerUid string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE customer_uid = $1", constants.CustomerTable)

	_, err := r.db.Exec(query, customerUid)
	if err != nil {
		r.log.Error("Error deleting customer", zap.Error(err))
		return err
	}

	return nil
}
