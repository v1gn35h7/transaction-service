package service

import "github.com/v1gn35h7/transaction-service/internal/models"

type AccountService interface {
	CreateAccount(document_number int64) (int64, error)
	GetAccount(accountID int64) (models.Account, error)
}

func (s service) CreateAccount(document_number int64) (int64, error) {
	accountID, err := s.ds.NewAccount(document_number)
	if err != nil {
		return 0, err
	}
	return accountID, nil
}

func (s service) GetAccount(accountID int64) (models.Account, error) {
	account, err := s.ds.GetAccount(accountID)
	if err != nil {
		return models.Account{}, err
	}
	return account, nil
}
