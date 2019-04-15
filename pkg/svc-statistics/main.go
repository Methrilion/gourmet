package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/methrilion/gourmet/pkg/util/connect"
)

type envStorage struct {
	gormDB *gorm.DB
}

var env envStorage

func main() {

	database := connect.GormDBConnect(os.Getenv("DB_DIALECT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSL"))

	env = envStorage{gormDB: database}

	fmt.Println("Hello, 世界")
}
