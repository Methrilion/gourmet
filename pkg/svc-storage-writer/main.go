package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/db"
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/mapper"
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/model"
	pbm "github.com/methrilion/gourmet/proto/model"
	pb "github.com/methrilion/gourmet/proto/svc-storage-writer/writer"
)

type storageWriterService struct {
	gormDB *gorm.DB
}

var storageWriter storageWriterService

func main() {

	keks := db.Get(os.Getenv("DB_DIALECT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSL"))

	storageWriter = storageWriterService{gormDB: keks}

	// Test functionality
	fmt.Println(storageWriter.ListCurrency(nil, &pb.ListCurrencyRequest{})) // Get all

	tes := pbm.Currency{Name: "kekekeke", Code: "123"}
	tesreq := &pb.CreateCurrencyRequest{Payload: &tes}
	fmt.Println(storageWriter.CreateCurrency(nil, tesreq)) // Create new

	fmt.Println("Hello, currency")
	cs := []model.Currency{}
	storageWriter.gormDB.Find(&cs) // Get all again
	fmt.Println(cs)
	// End test
}

func (s *storageWriterService) ListCurrency(ctx context.Context, in *pb.ListCurrencyRequest) (*pb.ListCurrencyResponse, error) {
	if in == nil {
		return &pb.ListCurrencyResponse{}, errors.New("Got nil request")
	}

	ms := []model.Currency{}
	log.Println("Get currency list:")

	if err := storageWriter.gormDB.Find(&ms).Error; err != nil {
		return nil, err
	}

	out := &pb.ListCurrencyResponse{}
	for _, r := range ms {
		out.Currency = append(out.Currency, mapper.CurrencyToPB(&r))
		fmt.Println(r)
	}

	return out, nil
}

func (s *storageWriterService) CreateCurrency(ctx context.Context, in *pb.CreateCurrencyRequest) (*pbm.Currency, error) {
	if in == nil {
		return &pbm.Currency{}, errors.New("Got nil request")
	}

	kek := mapper.PBToCurrency(in.GetPayload())
	m := model.Currency{}

	if storageWriter.gormDB.NewRecord(kek) != true {
		return nil, errors.New("Primary key is not blank before create")
	}
	if err := storageWriter.gormDB.Create(&kek).Scan(&m).Error; err != nil {
		return nil, err
	}
	storageWriter.gormDB.NewRecord(kek)
	if storageWriter.gormDB.NewRecord(kek) != false {
		return nil, errors.New("Primary key is blank after create")
	}

	return mapper.CurrencyToPB(&m), nil
}
