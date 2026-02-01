package tables

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(Up0002, Down0002)
}

func Up0002(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS accounts (
		account_id SERIAL PRIMARY KEY,
		document_number BIGINT NOT NULL
	);
	`)
	return err
}

func Down0002(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DROP TABLE IF EXISTS accounts;`)
	return err
}
