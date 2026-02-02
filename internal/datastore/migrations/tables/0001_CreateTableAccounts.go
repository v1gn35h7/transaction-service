package tables

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(Up0001, Down0001)
}

func Up0001(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS accounts (
		account_id SERIAL PRIMARY KEY,
		document_number BIGINT NOT NULL UNIQUE 
	);
	`)
	return err
}

func Down0001(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DROP TABLE IF EXISTS accounts;`)
	return err
}
