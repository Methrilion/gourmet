package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func Connect(dialect, host, port, user, dbname, password, sslmode string) *gorm.DB {
	log.Println("Connecting to Db...")

	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			host, port, user, dbname, password, sslmode)

	db, err := gorm.Open(dialect, connectionString)
	if err != nil {
		log.Fatal("Fatal:", err) // TODO: Fatal or Panic? Fatal crashes the service
	}

	log.Println("Connection successful")
	return db
}
