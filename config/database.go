package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/xvbnm48/go-network-media/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func RunMigrations(db *gorm.DB) error {
	// Definisi migrasi di sini
	err := db.AutoMigrate(&model.Reaction{})
	if err != nil {
		return err
	}

	return nil
}
