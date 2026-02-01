package config

import (
	"fmt"

	"github.com/go-logr/zerologr"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/transaction-service/internal/constants"
)

type PostgresqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

/*
* Reads config
 */
func ReadConfig(configPath string, logger zerologr.Logger) {
	// Read config
	logger.Info("Reading config from file", "confi_path", configPath)
	viper.Reset()
	viper.SetConfigName(constants.ConfigName) // name of config file (without extension)
	viper.SetConfigType(constants.ConfigType) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(configPath)           // path to look for the config file in
	viper.AddConfigPath(".")                  // optionally look for config in the working directory
	err := viper.ReadInConfig()               // Find and read the config file

	if err != nil {
		// Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// Prints all config
	//fmt.Println(viper.AllSettings())
}

func LoadPostgresqlConfig() *PostgresqlConfig {
	return &PostgresqlConfig{
		Host:     viper.GetString("postgresql.host"),
		Port:     viper.GetInt("postgresql.port"),
		User:     viper.GetString("postgresql.username"),
		Password: viper.GetString("postgresql.password"),
		DBName:   viper.GetString("postgresql.dbname"),
		SSLMode:  viper.GetString("postgresql.sslmode"),
	}
}
