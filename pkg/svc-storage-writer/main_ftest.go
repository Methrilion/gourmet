package main

import (
	"log"

	"github.com/methrilion/gourmet/pkg/svc-storage-writer/model"
	pbm "github.com/methrilion/gourmet/proto/model"
	pb "github.com/methrilion/gourmet/proto/svc-storage-writer/writer"
)

func AllTests() {
	fTestListCurrency()
	fTestCreateCurrency() // Create new
	fTestListCurrency2()  // Get all again
	fTestListCurrency3()
}

func fTestListCurrency() {
	log.Println(storageWriter.ListCurrency(nil, &pb.ListCurrencyRequest{}))
}

func fTestListCurrency3() {
	log.Println(storageWriter.ListCurrency(nil, &pb.ListCurrencyRequest{}))
}

func fTestCreateCurrency() {

	tes := pbm.Currency{Name: "testName", Code: "123"}
	tesreq := &pb.CreateCurrencyRequest{Payload: &tes}

	log.Println("fTestCreateCurrency() returns:")
	log.Println(storageWriter.CreateCurrency(nil, tesreq))
}

func fTestListCurrency2() {

	log.Println("func fTestListCurrency2()")
	cs := []model.Currency{}
	storageWriter.gormDB.Find(&cs)
	log.Println(cs)
}
