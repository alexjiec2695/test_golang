package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"test/internal/repository/postgres/data"
)

func NewConnection() (*gorm.DB, error) {
	dsn := os.Getenv("urlconnection")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(data.User{})
	db.AutoMigrate(data.Drug{})
	db.AutoMigrate(data.Vaccination{})
	return db, nil
}
