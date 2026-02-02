package tables

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(Up002, Down002)
}

func Up002(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS transactions (
		transaction_id SERIAL PRIMARY KEY,
		account_id BIGINT NOT NULL,
		operation_type_id INT NOT NULL,	
		amount NUMERIC(15,2) NOT NULL,
		event_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (account_id) REFERENCES accounts(account_id)
	);
	`)
	return err
}

func Down002(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DROP TABLE IF EXISTS transactions;`)
	return err
}
