package database

import (
	"app/middleware"
	"app/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error

	if err != nil {
		panic("failed to parse database port")
	}

	dsn := "host=localhost user=postgres password=root dbname=strix port=5437 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	err1 := db.AutoMigrate(&model.Admin{})
	if err1 != nil {
		fmt.Println("Can't migrate Product")
	}

	err3 := db.AutoMigrate(&model.Note{})
	if err3 != nil {
		fmt.Println("Can't migrate Note")
	}

	err2 := db.AutoMigrate(&model.User{})
	if err2 != nil {
		fmt.Println("Can't migrate User")
	}

	err4 := db.AutoMigrate(&model.Group{})
	if err4 != nil {
		fmt.Println("Can't migrate Group")
	}

	fmt.Println("Database Migrated")

	password, _ := middleware.HashPassword("root1234")
	db.Create(&model.Admin{Username: "Admin", Email: "admin@admin.ru", Password: password})

	DB = db
}
