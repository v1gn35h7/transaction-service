package postgresql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "postgres")
	ds := &Datastore{DB: sqlxDB}
	documentNumber := int64(123456789)
	accountID := int64(1)
	mock.ExpectQuery(`INSERT INTO accounts \(document_number\) VALUES \(\$1\) RETURNING account_id;`).
		WithArgs(documentNumber).
		WillReturnRows(sqlmock.NewRows([]string{"account_id"}).AddRow(accountID))
	returnedAccountID, err := ds.NewAccount(documentNumber)
	if err != nil {
		t.Fatalf("failed to create account: %v", err)
	}
	assert.Equal(t, accountID, returnedAccountID, "expected account_id %d, got %d", accountID, returnedAccountID)
}

func TestNewAccount_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "postgres")
	ds := &Datastore{DB: sqlxDB}
	documentNumber := int64(123456789)
	mock.ExpectQuery(`INSERT INTO accounts \(document_number\) VALUES \(\$1\) RETURNING account_id;`).
		WithArgs(documentNumber).
		WillReturnError(assert.AnError)
	_, err = ds.NewAccount(documentNumber)
	if err == nil {
		t.Fatalf("expected error but got none")
	}
	assert.Equal(t, assert.AnError, err, "expected error %v, got %v", assert.AnError, err)
}

func TestGetAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "postgres")
	ds := &Datastore{DB: sqlxDB}
	accountID := int64(1)
	documentNumber := int64(123456789)
	mock.ExpectQuery(`SELECT account_id, document_number FROM accounts WHERE account_id = \$1;`).
		WithArgs(accountID).
		WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number"}).
			AddRow(accountID, documentNumber))
	account, err := ds.GetAccount(accountID)
	if err != nil {
		t.Fatalf("failed to get account: %v", err)
	}
	assert.Equal(t, accountID, account.AccountID, "expected account_id %d, got %d", accountID, account.AccountID)
	assert.Equal(t, documentNumber, account.DocumentNumber, "expected document_number %d, got %d", documentNumber, account.DocumentNumber)
}

func TestGetAccount_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "postgres")
	ds := &Datastore{DB: sqlxDB}
	accountID := int64(1)
	mock.ExpectQuery(`SELECT account_id, document_number FROM accounts WHERE account_id = \$1;`).
		WithArgs(accountID).
		WillReturnError(assert.AnError)
	_, err = ds.GetAccount(accountID)
	if err == nil {
		t.Fatalf("expected error but got none")
	}
	assert.Equal(t, assert.AnError, err, "expected error %v, got %v", assert.AnError, err)
}
