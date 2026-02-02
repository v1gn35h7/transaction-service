package models

type Transaction struct {
	AccountID       int64   `json:"account_id"`
	TransactionID   int64   `json:"transaction_id"`
	Amount          float64 `json:"amount"`
	OperationTypeID int16   `json:"transaction_type_id"`
	EventDate       string  `json:"event_date"`
}

type TransactionsDatastore interface {
	NewTransaction(accountID int64, operationTypeID int16, amount float64) (*Transaction, error)
	ListTransactions(accountID int64) ([]Transaction, error)
}
