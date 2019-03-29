package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/db"
	//pb "github.com/methrilion/gourmet/proto/svc-storage-writer"
)

// Currency model golang implementation
type Currency struct {
	ID   uint32
	Name string
	Code string
}

// TableName set Currency's table name to be `currency`
func (Currency) TableName() string {
	return "currency"
}

func main() {
	fmt.Println("Hello, currency")

	cs := []Currency{}
	db.Get().Find(&cs)

	fmt.Println(cs)
}
