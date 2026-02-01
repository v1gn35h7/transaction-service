package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/v1gn35h7/transaction-service/internal/logging"
)

func TestInvalidConfigPath(t *testing.T) {
	logger := logging.Logger()
	invalidConfigPath := "/opt/"
	t.Run("Read invalid config file", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic due to invalid config path, but got none")
			}
		}()
		ReadConfig(invalidConfigPath, logger)
	})
}

func TestReadConfig(t *testing.T) {
	logger := logging.Logger()
	configPath := "../../"

	t.Run("Read valid config file", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Expected no panic, but got: %v", r)
			}
		}()
		ReadConfig(configPath, logger)
	})

}

func TestLoadPostgresqlConfig(t *testing.T) {
	logger := logging.Logger()
	configPath := "../../"
	t.Run("Load Postgresql Config", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Expected no panic, but got: %v", r)
			}
		}()
		ReadConfig(configPath, logger)
		pgConfig := LoadPostgresqlConfig()
		assert.NotNil(t, pgConfig, "PostgresqlConfig should not be nil")
	})
}
