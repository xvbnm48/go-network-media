package config

import "os"

type databaseConfig struct{}

func NewDatabaseConfig() DatabaseConfig {
	return &databaseConfig{}
}

func (d *databaseConfig) GetUsername() string {
	return os.Getenv("DB_USER")
}

func (d *databaseConfig) GetPassword() string {
	return os.Getenv("DB_PASS")
}

func (d *databaseConfig) GetHost() string {
	return os.Getenv("DB_HOST")
}

func (d *databaseConfig) GetPort() string {
	return os.Getenv("DB_PORT")
}

func (d *databaseConfig) GetDatabaseName() string {
	return os.Getenv("DB_NAME")
}
