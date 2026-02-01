package postgresql

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/v1gn35h7/transaction-service/internal/config"
	"github.com/v1gn35h7/transaction-service/internal/logging"
)

func TestGenerateDNS(t *testing.T) {
	conf := &config.PostgresqlConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "user",
		Password: "password",
		DBName:   "testdb",
		SSLMode:  "disable",
	}

	t.Run("Valid config", func(t *testing.T) {
		dns := generateDSN(conf)
		assert.Equal(t, "host=localhost port=5432 user=user password=password dbname=testdb sslmode=disable", dns, "DSN should match expected format")
	})

	confMissing := &config.PostgresqlConfig{
		Host: "localhost",
		Port: 5432,
	}
	t.Run("Missing fields in config", func(t *testing.T) {
		dns := generateDSN(confMissing)
		assert.Equal(t, "host=localhost port=5432 user= password= dbname= sslmode=", dns, "DSN should handle missing fields")
	})
}

func TestNewDatastore_Success(t *testing.T) {
	old := sqlxConnect
	defer func() { sqlxConnect = old }()

	sqlDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	sqlxConnect = func(driverName, dataSourceName string) (*sqlx.DB, error) {
		return sqlx.NewDb(sqlDB, "postgres"), nil
	}

	logger := logging.Logger()
	conf := &config.PostgresqlConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "user",
		Password: "password",
		DBName:   "testdb",
		SSLMode:  "disable",
	}

	ds, err := NewDatastore(conf, logger)
	assert.NoError(t, err)
	assert.NotNil(t, ds)
	if ds != nil && ds.DB != nil {
		ds.DB.Close()
	}
}

func TestNewDatastore_PingFailure(t *testing.T) {
	old := sqlxConnect
	defer func() { sqlxConnect = old }()

	sqlDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	// Expect ping to fail
	mock.ExpectPing().WillReturnError(fmt.Errorf("ping failed"))

	sqlxConnect = func(driverName, dataSourceName string) (*sqlx.DB, error) {
		return sqlx.NewDb(sqlDB, "postgres"), nil
	}

	logger := logging.Logger()
	conf := &config.PostgresqlConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "user",
		Password: "password",
		DBName:   "testdb",
		SSLMode:  "disable",
	}

	ds, err := NewDatastore(conf, logger)
	assert.Error(t, err)
	assert.Nil(t, ds)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}
