package models

type Account struct {
	AccountID      int64 `json:"account_id"`
	DocumentNumber int64 `json:"document_number"`
}

type AccountsDatastore interface {
	NewAccount(documentNumber int64) (int64, error)
	GetAccount(accountID int64) (Account, error)
}
