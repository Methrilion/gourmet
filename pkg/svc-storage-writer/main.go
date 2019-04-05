package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/db"
	pbm "github.com/methrilion/gourmet/proto/model"
	pb "github.com/methrilion/gourmet/proto/writer"
	"google.golang.org/grpc"
)

type storageWriterService struct {
	gormDB *gorm.DB
}

var storageWriter storageWriterService

func main() {
	time.Sleep(2 * time.Second)

	database := db.Connect(os.Getenv("DB_DIALECT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSL"))

	storageWriter = storageWriterService{gormDB: database}

	lis, err := net.Listen("tcp", os.Getenv("STORAGE_WRITER_SERVICE_ADDRESS"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	pb.RegisterStorageWriterServiceServer(s, &storageWriter)

	log.Println("Now listening on", os.Getenv("STORAGE_WRITER_SERVICE_ADDRESS"))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
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

func (s *storageWriterService) ListRatesOfExchange(ctx context.Context, in *pb.ListRatesOfExchangeRequest) (*pb.ListRatesOfExchangeResponse, error) {

	result, err := pbm.DefaultListRateOfExchange(ctx, storageWriter.gormDB)

	return &pb.ListRatesOfExchangeResponse{
		RatesOfExchange: result,
	}, err
}

func (s *storageWriterService) CreateRateOfExchange(ctx context.Context, in *pb.CreateRateOfExchangeRequest) (*pbm.RateOfExchange, error) {

	return pbm.DefaultCreateRateOfExchange(ctx, in.GetPayload(), storageWriter.gormDB)
}

func (s *storageWriterService) ListLocations(ctx context.Context, in *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {

	result, err := pbm.DefaultListLocation(ctx, storageWriter.gormDB)

	return &pb.ListLocationsResponse{
		Locations: result,
	}, err
}

func (s *storageWriterService) CreateLocation(ctx context.Context, in *pb.CreateLocationRequest) (*pbm.Location, error) {

	return pbm.DefaultCreateLocation(ctx, in.GetPayload(), storageWriter.gormDB)
}

func (s *storageWriterService) ListProducts(ctx context.Context, in *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {

	result, err := pbm.DefaultListProduct(ctx, storageWriter.gormDB)

	return &pb.ListProductsResponse{
		Products: result,
	}, err
}

func (s *storageWriterService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pbm.Product, error) {

	return pbm.DefaultCreateProduct(ctx, in.GetPayload(), storageWriter.gormDB)
}

func (s *storageWriterService) ListPrices(ctx context.Context, in *pb.ListPricesRequest) (*pb.ListPricesResponse, error) {

	result, err := pbm.DefaultListPrice(ctx, storageWriter.gormDB)

	return &pb.ListPricesResponse{
		Prices: result,
	}, err
}

func (s *storageWriterService) CreatePrice(ctx context.Context, in *pb.CreatePriceRequest) (*pbm.Price, error) {

	return pbm.DefaultCreatePrice(ctx, in.GetPayload(), storageWriter.gormDB)
}

func (s *storageWriterService) ListPositions(ctx context.Context, in *pb.ListPositionsRequest) (*pb.ListPositionsResponse, error) {

	result, err := pbm.DefaultListPosition(ctx, storageWriter.gormDB)

	return &pb.ListPositionsResponse{
		Positions: result,
	}, err
}

func (s *storageWriterService) CreatePosition(ctx context.Context, in *pb.CreatePositionRequest) (*pbm.Position, error) {

	return pbm.DefaultCreatePosition(ctx, in.GetPayload(), storageWriter.gormDB)
}

func (s *storageWriterService) ListEmployees(ctx context.Context, in *pb.ListEmployeesRequest) (*pb.ListEmployeesResponse, error) {

	result, err := pbm.DefaultListEmployee(ctx, storageWriter.gormDB)

	return &pb.ListEmployeesResponse{
		Employees: result,
	}, err
}

func (s *storageWriterService) CreateEmployee(ctx context.Context, in *pb.CreateEmployeeRequest) (*pbm.Employee, error) {

	return pbm.DefaultCreateEmployee(ctx, in.GetPayload(), storageWriter.gormDB)
}

func (s *storageWriterService) ListMethods(ctx context.Context, in *pb.ListMethodsRequest) (*pb.ListMethodsResponse, error) {

	result, err := pbm.DefaultListMethod(ctx, storageWriter.gormDB)

	return &pb.ListMethodsResponse{
		Methods: result,
	}, err
}

func (s *storageWriterService) CreateMethod(ctx context.Context, in *pb.CreateMethodRequest) (*pbm.Method, error) {

	return pbm.DefaultCreateMethod(ctx, in.GetPayload(), storageWriter.gormDB)
}

func (s *storageWriterService) ListReceipts(ctx context.Context, in *pb.ListReceiptsRequest) (*pb.ListReceiptsResponse, error) {

	result, err := pbm.DefaultListReceipt(ctx, storageWriter.gormDB)

	return &pb.ListReceiptsResponse{
		Receipts: result,
	}, err
}

func (s *storageWriterService) CreateReceipt(ctx context.Context, in *pb.CreateReceiptRequest) (*pbm.Receipt, error) {

	return pbm.DefaultCreateReceipt(ctx, in.GetPayload(), storageWriter.gormDB)
}

func (s *storageWriterService) ListPurchases(ctx context.Context, in *pb.ListPurchasesRequest) (*pb.ListPurchasesResponse, error) {

	result, err := pbm.DefaultListPurchase(ctx, storageWriter.gormDB)

	return &pb.ListPurchasesResponse{
		Purchases: result,
	}, err
}

func (s *storageWriterService) CreatePurchase(ctx context.Context, in *pb.CreatePurchaseRequest) (*pbm.Purchase, error) {

	return pbm.DefaultCreatePurchase(ctx, in.GetPayload(), storageWriter.gormDB)
}
