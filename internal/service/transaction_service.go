package service

import "github.com/v1gn35h7/transaction-service/internal/models"

type TransactionService interface {
	CreateTransaction(account_id int64, operation_type_id int16, amount float64) error
	ListTransactions(account_id int64) ([]models.Transaction, error)
}

func (s service) CreateTransaction(account_id int64, operation_type_id int16, amount float64) error {
	_, err := s.ds.NewTransaction(account_id, operation_type_id, amount)
	if err != nil {
		return err
	}
	return nil
}

func (s service) ListTransactions(account_id int64) ([]models.Transaction, error) {
	transactions, err := s.ds.ListTransactions(account_id)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
