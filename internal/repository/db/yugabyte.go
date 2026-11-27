package db

import (
	"fmt"
	"log"
	pgdriver "gorm.io/driver/postgres"
	"prob13/internal/core/entity"

	"gorm.io/gorm"
)

func ConnectYugabyte() *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s",
		"localhost", "5436", "yugabyte", "yugabyte", "disable",
	)

	db, err := gorm.Open(pgdriver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to yugabyte")
	}

	db.AutoMigrate(&entity.User{}, &entity.UserGroup{})
	fmt.Println("Yugabyte connected successfully")
	return db
}
