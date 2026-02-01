package postgresql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "postgres")
	ds := &Datastore{DB: sqlxDB}
	accountID := int64(1)
	operationTypeID := int16(1)
	amount := 100.50
	transactionID := int64(1)
	eventDate := "2024-01-01T00:00:00Z"
	mock.ExpectQuery(`INSERT INTO transactions \(account_id, operation_type_id, amount\) VALUES \(\$1, \$2, \$3\) RETURNING transaction_id, account_id, operation_type_id, amount, event_date;`).
		WithArgs(accountID, operationTypeID, amount).
		WillReturnRows(sqlmock.NewRows([]string{"transaction_id", "account_id", "operation_type_id", "amount", "event_date"}).
			AddRow(transactionID, accountID, operationTypeID, amount, eventDate))
	returnedTransaction, err := ds.NewTransaction(accountID, operationTypeID, amount)
	if err != nil {
		t.Fatalf("failed to create transaction: %v", err)
	}
	assert.Equal(t, transactionID, returnedTransaction.TransactionID, "expected transaction_id %d, got %d", transactionID, returnedTransaction.TransactionID)
	assert.Equal(t, accountID, returnedTransaction.AccountID, "expected account_id %d, got %d", accountID, returnedTransaction.AccountID)
	assert.Equal(t, operationTypeID, returnedTransaction.OperationTypeID, "expected operation_type_id %d, got %d", operationTypeID, returnedTransaction.OperationTypeID)
	assert.Equal(t, amount, returnedTransaction.Amount, "expected amount %f, got %f", amount, returnedTransaction.Amount)
	assert.Equal(t, eventDate, returnedTransaction.EventDate, "expected event_date %s, got %s", eventDate, returnedTransaction.EventDate)
}

func TestNewTransaction_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "postgres")
	ds := &Datastore{DB: sqlxDB}
	accountID := int64(1)
	operationTypeID := int16(1)
	amount := 100.50
	mock.ExpectQuery(`INSERT INTO transactions \(account_id, operation_type_id, amount\) VALUES \(\$1, \$2, \$3\) RETURNING transaction_id, account_id, operation_type_id, amount, event_date;`).
		WithArgs(accountID, operationTypeID, amount).
		WillReturnError(assert.AnError)
	_, err = ds.NewTransaction(accountID, operationTypeID, amount)
	if err == nil {
		t.Fatalf("expected error but got none")
	}
	assert.Equal(t, assert.AnError, err, "expected error %v, got %v", assert.AnError, err)
}
