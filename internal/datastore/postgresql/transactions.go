package postgresql

import "github.com/v1gn35h7/transaction-service/internal/models"

func (ds *Datastore) NewTransaction(account_id int64, operation_type_id int16, amount float64) (*models.Transaction, error) {
	var transaction models.Transaction
	sqlStatement := `INSERT INTO transactions (account_id, operation_type_id, amount) VALUES ($1, $2, $3) RETURNING transaction_id, account_id, operation_type_id, amount, event_date;`
	err := ds.DB.QueryRow(sqlStatement, account_id, operation_type_id, amount).Scan(&transaction.TransactionID, &transaction.AccountID, &transaction.OperationTypeID, &transaction.Amount, &transaction.EventDate)
	return &transaction, err
}
