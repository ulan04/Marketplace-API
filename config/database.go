package config

import (
	"fmt"
	"log"
	"marketplace-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=123456 dbname=marketplace_db port=5432 sslmode=disable TimeZone=Asia/Almaty"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Database connected successfully")

	err = db.AutoMigrate(&models.User{}, &models.Category{}, &models.Item{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	fmt.Println("Database migrated successfully")

	DB = db
}