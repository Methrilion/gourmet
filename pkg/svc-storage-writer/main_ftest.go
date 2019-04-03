package main

import (
	"fmt"

	"github.com/methrilion/gourmet/pkg/svc-storage-writer/model"
	pbm "github.com/methrilion/gourmet/proto/model"
	pb "github.com/methrilion/gourmet/proto/svc-storage-writer/writer"
)

func AllTests() {
	fTestListCurrency()
	fTestCreateCurrencyWithNilRequest()
	fTestCreateCurrency() // Create new
	fTestListCurrency2()  // Get all again
}

func fTestListCurrency() {

	fmt.Println(storageWriter.ListCurrency(nil, &pb.ListCurrencyRequest{}))
}

func fTestCreateCurrencyWithNilRequest() {

	fmt.Println(storageWriter.CreateCurrency(nil, nil))
}

func fTestCreateCurrency() {

	tes := pbm.Currency{Name: "testName", Code: "123"}
	tesreq := &pb.CreateCurrencyRequest{Payload: &tes}

	fmt.Println("fTestCreateCurrency() returns:")
	fmt.Println(storageWriter.CreateCurrency(nil, tesreq))
}

func fTestListCurrency2() {

	fmt.Println("func fTestListCurrency2()")
	cs := []model.Currency{}
	storageWriter.gormDB.Find(&cs)
	fmt.Println(cs)
}
