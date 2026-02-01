package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/v1gn35h7/transaction-service/internal/config"
	"github.com/v1gn35h7/transaction-service/internal/datastore/migrations/tables"
)

var sqlxConnect = sqlx.Connect

const maxAttempts = 5

type Datastore struct {
	DB     *sqlx.DB
	logger logr.Logger
}

func NewDatastore(conf *config.PostgresqlConfig, logger logr.Logger) (*Datastore, error) {
	dsn := generateDSN(conf)
	logger.Info("Connecting to Postgresql", "dsn", dsn)
	db, err := sqlxConnect("postgres", dsn)
	if err != nil {
		logger.Error(err, "Failed to connect to Postgresql")
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(0)

	var dbError error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		dbError = db.Ping()
		if dbError == nil {
			// we're connected!
			break
		}
		interval := time.Duration(attempt) * time.Second
		logger.Info("Postgresql: could not connect to db", "error", dbError, "sleeping", interval)
		time.Sleep(interval)
	}

	if dbError != nil {
		return nil, dbError
	}

	return &Datastore{
		DB:     db,
		logger: logger,
	}, nil
}

func generateDSN(conf *config.PostgresqlConfig) string {
	dns := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	return fmt.Sprintf(dns, conf.Host, conf.Port, conf.User, conf.Password, conf.DBName, conf.SSLMode)
}

func (ds *Datastore) Close() error {
	return ds.DB.Close()
}

func (ds *Datastore) RunMigrations() error {
	ds.logger.Info("Running database migrations...")
	client := tables.NewMigrationClient(ds.DB.DB)
	ctx := context.Background()
	_, err := client.Up(ctx)
	if err != nil {
		return err
	}
	return nil
}
