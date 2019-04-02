package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func connect(dialect, host, port, user, dbname, password, sslmode string) *gorm.DB {
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

//Get - kek realization
func Get(dialect, host, port, user, dbname, password, sslmode string) *gorm.DB {

	// TODO: Add Get() functionality. Now it just a plug
	db := connect(dialect, host, port, user, dbname, password, sslmode)
	return db
}
