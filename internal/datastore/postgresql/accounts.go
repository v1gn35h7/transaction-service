package postgresql

import "github.com/v1gn35h7/transaction-service/internal/models"

func (ds *Datastore) NewAccount(documentNumber int64) (int64, error) {
	var accountID int64
	sqlStatement := `INSERT INTO accounts (document_number) VALUES ($1) RETURNING account_id;`
	err := ds.DB.QueryRow(sqlStatement, documentNumber).Scan(&accountID)
	if err != nil {
		return 0, err
	}
	return accountID, nil
}

func (ds *Datastore) GetAccount(accountID int64) (models.Account, error) {
	var account models.Account
	sqlStatement := `SELECT account_id, document_number FROM accounts WHERE account_id = $1;`
	err := ds.DB.QueryRow(sqlStatement, accountID).Scan(&account.AccountID, &account.DocumentNumber)
	if err != nil {
		return models.Account{}, err
	}
	return account, nil
}
