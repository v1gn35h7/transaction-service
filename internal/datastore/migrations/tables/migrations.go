package tables

import (
	"database/sql"
	"os"

	"github.com/pressly/goose/v3"
)

var (
	MigrationClient *goose.Provider
)

func NewMigrationClient(db *sql.DB) *goose.Provider {
	if MigrationClient == nil {
		MigrationClient, _ = goose.NewProvider(goose.DialectPostgres, db, os.DirFS("internal/datastore/migrations/tables"))
	}

	return MigrationClient
}
