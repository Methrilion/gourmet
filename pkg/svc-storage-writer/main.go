package main

import (
	"context"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/db"
	pbm "github.com/methrilion/gourmet/proto/model"
	pb "github.com/methrilion/gourmet/proto/svc-storage-writer/writer"
)

type storageWriterService struct {
	gormDB *gorm.DB
}

var storageWriter storageWriterService

func main() {

	database := db.Connect(os.Getenv("DB_DIALECT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSL"))

	storageWriter = storageWriterService{gormDB: database}

	AllTests()
}

func (s *storageWriterService) ListCurrency(ctx context.Context, in *pb.ListCurrencyRequest) (*pb.ListCurrencyResponse, error) {

	result, err := pbm.DefaultListCurrency(ctx, storageWriter.gormDB)

	return &pb.ListCurrencyResponse{
		Currency: result,
	}, err
}

func (s *storageWriterService) CreateCurrency(ctx context.Context, in *pb.CreateCurrencyRequest) (*pbm.Currency, error) {

	return pbm.DefaultCreateCurrency(ctx, in.GetPayload(), storageWriter.gormDB)
}

// func (s *storageWriterService) ListRatesOfExchange(ctx context.Context, in *pb.ListRatesOfExchangeRequest) (*pb.ListRatesOfExchangeResponse, error)
// func (s *storageWriterService) CreateRateOfExchange(ctx context.Context, in *pb.CreateRateOfExchangeRequest) (*pbm.RateOfExchange, error)
// func (s *storageWriterService) ListLocations(ctx context.Context, in *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error)
// func (s *storageWriterService) CreateLocadb.Get()tion(ctx context.Context, in *pb.CreateLocationRequest) (*pbm.Location, error)
// func (s *storageWriterService) ListProducts(ctx context.Context, in *pb.ListProductsRequest) (*pb.ListProductsResponse, error)
// func (s *storageWriterService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pbm.Product, error)
// func (s *storageWriterService) ListPrices(ctx context.Context, in *pb.ListPricesRequest) (*pb.ListPricesResponse, error)
// func (s *storageWriterService) CreatePrice(ctx context.Context, in *pb.CreatePriceRequest) (*pbm.Price, error)
// func (s *storageWriterService) ListPositions(ctx context.Context, in *pb.ListPositionsRequest) (*pb.ListPositionsResponse, error)
// func (s *storageWriterService) CreatePosition(ctx context.Context, in *pb.CreatePositionRequest) (*pbm.Position, error)
// func (s *storageWriterService) ListEmployees(ctx context.Context, in *pb.ListEmployeesRequest) (*pb.ListEmployeesResponse, error)
// func (s *storageWriterService) CreateEmployee(ctx context.Context, in *pb.CreateEmployeeRequest) (*pbm.Employee, error)
// func (s *storageWriterService) ListMethods(ctx context.Context, in *pb.ListMethodsRequest) (*pb.ListMethodsResponse, error)
// func (s *storageWriterService) CreateMethod(ctx context.Context, in *pb.CreateMethodRequest) (*pbm.Method, error)
// func (s *storageWriterService) ListReceipts(ctx context.Context, in *pb.ListReceiptsRequest) (*pb.ListReceiptsResponse, error)
// func (s *storageWriterService) CreateReceipt(ctx context.Context, in *pb.CreateReceiptRequest) (*pbm.Receipt, error)
// func (s *storageWriterService) ListPurchases(ctx context.Context, in *pb.ListPurchasesRequest) (*pb.ListPurchasesResponse, error)
// func (s *storageWriterService) CreatePurchase(ctx context.Context, in *pb.CreatePurchaseRequest) (*pbm.Purchase, error)
