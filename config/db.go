package config

type DatabaseConfig interface {
	GetUsername() string
	GetPassword() string
	GetHost() string
	GetPort() string
	GetDatabaseName() string
}
