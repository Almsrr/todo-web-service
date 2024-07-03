package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


var DB *sql.DB

func ConnectToDb() {
	env, _ := godotenv.Read(".env")

	config := mysql.Config{
		User:   env["DBUSER"],
		Passwd: env["DBPASS"],
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "Todos-Service",
		AllowNativePasswords: true,
	}	

	var err error
	DB, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to database")
}
