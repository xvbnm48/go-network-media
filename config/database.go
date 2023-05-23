package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetUpDatabase() (*gorm.DB, error) {
	err := godotenv.Load(".env") // Load file .env
	if err != nil {
		return nil, err
	}

	dbConfig := NewDatabaseConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.GetUsername(), dbConfig.GetPassword(), dbConfig.GetHost(),
		dbConfig.GetPort(), dbConfig.GetDatabaseName())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
