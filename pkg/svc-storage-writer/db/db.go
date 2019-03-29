package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {

	db = connect()
}

func connect() *gorm.DB {
	log.Println("Connecting to Db...")

	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_DBNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_SSL"))

	db, err := gorm.Open(os.Getenv("DB_DIALECT"), connectionString)
	if err != nil {
		log.Fatal("Fatal:", err)
	}

	log.Println("Connection successful")
	return db
}

//Get - Singleton realization
func Get() *gorm.DB {
	if db != nil {
		return db
	}
	db := connect()
	return db
}
