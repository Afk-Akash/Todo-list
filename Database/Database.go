package database

import (
	"fmt"
	"log"
	"os"
	models "todolist/Models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	DSN := os.Getenv("DSN")

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("error occured while connecting to db \n", err)
	}
	fmt.Println("connected to database....")
	DB = db

	DB.AutoMigrate(
		&models.User{},
		&models.Todo{},
	)
}
