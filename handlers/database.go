package handlers

import (
	"almsrr/todo-web-service/models"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	env, _ := godotenv.Read(".env")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", env["DB_USER"], env["DB_PASS"], env["DB_HOST"], env["DB_PORT"], env["DB_NAME"])

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")

	createTables()
}

func createTables() {
	DB.AutoMigrate(&models.Todo{})

	log.Println("Tables created")
}
