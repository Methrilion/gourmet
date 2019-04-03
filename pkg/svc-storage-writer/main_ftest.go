package main

import (
	"log"
	"time"

	pbm "github.com/methrilion/gourmet/proto/model"
	pb "github.com/methrilion/gourmet/proto/svc-storage-writer/writer"
)

func AllTests() {
	time.Sleep(1 * time.Second)

	fTestListCurrency()
	fTestCreateCurrencyWithNilRequest()
	fTestCreateCurrency() // Create new
	fTestListCurrency2()  // Get all again
	fTestListProducts()
	fTestListROE()
	fTestCreateROE()
	fTestListROE()
}

func fTestListROE() {
	log.Println("fTestListROE():")
	log.Println(storageWriter.ListRatesOfExchange(nil, &pb.ListRatesOfExchangeRequest{}))
}

func fTestCreateROE() {
	log.Println("fTestCreateROE():")

	tes := pbm.RateOfExchange{FromId: 1, ToId: 2, Price: 330}
	tesreq := &pb.CreateRateOfExchangeRequest{Payload: &tes}

	log.Println(storageWriter.CreateRateOfExchange(nil, tesreq))
}

func fTestListProducts() {
	log.Println("fTestListProducts():")
	log.Println(storageWriter.ListProducts(nil, &pb.ListProductsRequest{}))
}

func fTestListCurrency() {
	log.Println("fTestListCurrency():")
	log.Println(storageWriter.ListCurrency(nil, &pb.ListCurrencyRequest{}))
}

func fTestCreateCurrencyWithNilRequest() {
	log.Println("fTestCreateCurrencyWithNilRequest():")
	log.Println(storageWriter.CreateCurrency(nil, nil))
}

func fTestCreateCurrency() {
	log.Println("fTestCreateCurrency():")

	tes := pbm.Currency{Name: "testName", Code: "123"}
	tesreq := &pb.CreateCurrencyRequest{Payload: &tes}

	log.Println(storageWriter.CreateCurrency(nil, tesreq))
}

func fTestListCurrency2() {
	log.Println("fTestListCurrency2():")
	cs := []pbm.CurrencyORM{}
	storageWriter.gormDB.Find(&cs)
	log.Println(cs)
}
