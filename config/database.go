package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT)
	db, errorDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect database")
	}

	return db
}

func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}
