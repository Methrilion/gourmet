package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/methrilion/gourmet/pkg/util/connect"
	pbm "github.com/methrilion/gourmet/proto/model"
	pb "github.com/methrilion/gourmet/proto/reader"
	"google.golang.org/grpc"
)

type storageReaderService struct {
	gormDB *gorm.DB
}

var storageReader storageReaderService

func main() {

	database := connect.GormDBConnect(os.Getenv("DB_DIALECT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSL"))

	storageReader = storageReaderService{gormDB: database}

	lis, err := net.Listen("tcp", os.Getenv("STORAGE_READER_SERVICE_ADDRESS"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	pb.RegisterStorageReaderServiceServer(s, &storageReader)

	log.Println("Now listening on", os.Getenv("STORAGE_READER_SERVICE_ADDRESS"))

	////////////////////// DELETE THIS PART ////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	time.Sleep(5 * time.Second)
	storageReader.gormDB.LogMode(true)
	storageReader.GetCountPurchasesByYear(nil,
		&pb.GetCountPurchasesByYearRequest{Year: 2016})

	start, _ := ptypes.TimestampProto(time.Date(2016, 7, 1, 0, 0, 0, 0, time.UTC))
	end, _ := ptypes.TimestampProto(time.Date(2016, 8, 1, 0, 0, 0, -1, time.UTC))
	storageReader.GetRevenuePurchasesByDatesByPrice(nil,
		&pb.GetRevenuePurchasesByDatesByPriceRequest{
			PriceId: 3,
			Start:   start,
			End:     end,
		})

	storageReader.GetRevenuePurchasesByDatesByProduct(nil,
		&pb.GetRevenuePurchasesByDatesByProductRequest{
			ProductId: 3,
			Start:     start,
			End:       end,
		})
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

func getMonths(year int) ([12]time.Time, [12]time.Time) {
	var starts [12]time.Time
	var ends [12]time.Time

	for i := 0; i < 12; i++ {
		starts[i] = time.Date(year, time.Month(i+1), 1, 0, 0, 0, 0, time.UTC)
		ends[i] = time.Date(year, time.Month(i+2), 1, 0, 0, 0, -1, time.UTC)
	}
	return starts, ends
}

func (s *storageReaderService) GetCountPurchasesByYear(ctx context.Context, in *pb.GetCountPurchasesByYearRequest) (*pb.GetCountPurchasesByYearResponse, error) {

	year := int(in.GetYear())

	type ResultQuery struct {
		Datetime time.Time
		Id       uint32
	}
	var resq []ResultQuery

	err := storageReader.gormDB.Table("purchases").
		Select("purchases.id, datetime").
		Joins("JOIN receipts ON purchases.receipt_id = receipts.id").
		Where("datetime BETWEEN ? AND ?",
			time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(year+1, 1, 1, 0, 0, 0, -1, time.UTC)).
		Scan(&resq).Error
	if err != nil {
		return nil, err
	}

	startMonths, endMonths := getMonths(year)
	var counts [12]int32

	for _, a := range resq {
		for i := 0; i < 12; i++ {
			if a.Datetime.After(startMonths[i]) && a.Datetime.Before(endMonths[i]) {
				counts[i]++
			}
		}
	}

	return &pb.GetCountPurchasesByYearResponse{
		Counts: counts[:],
	}, nil
}

func (s *storageReaderService) GetRevenuePurchasesByDatesByPrice(ctx context.Context, in *pb.GetRevenuePurchasesByDatesByPriceRequest) (*pb.GetRevenuePurchasesByDatesByPriceResponse, error) {

	start, err := ptypes.Timestamp(in.GetStart())
	if err != nil {
		return nil, err
	}

	end, err := ptypes.Timestamp(in.GetEnd())
	if err != nil {
		return nil, err
	}
	fmt.Println(start, end)

	type Result struct {
		Total float32
	}
	var revenue Result

	err = storageReader.gormDB.
		Raw("SELECT SUM(result) AS total "+
			"FROM purchases, receipts "+
			"WHERE purchases.price_id = ? "+
			"AND purchases.receipt_id = receipts.id "+
			"AND receipts.datetime BETWEEN ? AND ?",
			in.GetPriceId(), start, end).
		Scan(&revenue).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(revenue)

	return &pb.GetRevenuePurchasesByDatesByPriceResponse{
		Revenue: revenue.Total,
	}, nil
}

func (s *storageReaderService) GetRevenuePurchasesByDatesByProduct(ctx context.Context, in *pb.GetRevenuePurchasesByDatesByProductRequest) (*pb.GetRevenuePurchasesByDatesByProductResponse, error) {

	start, err := ptypes.Timestamp(in.GetStart())
	if err != nil {
		return nil, err
	}

	end, err := ptypes.Timestamp(in.GetEnd())
	if err != nil {
		return nil, err
	}

	type Result struct {
		Total float32
	}
	var revenue Result

	err = storageReader.gormDB.Table("purchases").
		Select("sum(result) as total").
		Joins("JOIN receipts ON purchases.receipt_id = receipts.id AND datetime BETWEEN ? AND ?", start, end).
		Joins("JOIN prices ON purchases.receipt_id = prices.id AND product_id = ?", in.GetProductId()).
		Scan(&revenue).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(revenue)

	return &pb.GetRevenuePurchasesByDatesByProductResponse{
		Revenue: revenue.Total,
	}, nil
}

func (s *storageReaderService) ListCurrency(ctx context.Context, in *pb.ListCurrencyRequest) (*pb.ListCurrencyResponse, error) {

	result, err := pbm.DefaultListCurrency(ctx, storageReader.gormDB)

	return &pb.ListCurrencyResponse{
		Currency: result,
	}, err
}

func (s *storageReaderService) GetCurrency(ctx context.Context, in *pb.GetCurrencyRequest) (*pbm.Currency, error) {

	return pbm.DefaultReadCurrency(ctx, &pbm.Currency{Id: in.GetId()}, storageReader.gormDB)
}

func (s *storageReaderService) ListRatesOfExchange(ctx context.Context, in *pb.ListRatesOfExchangeRequest) (*pb.ListRatesOfExchangeResponse, error) {

	result, err := pbm.DefaultListRateOfExchange(ctx, storageReader.gormDB)

	return &pb.ListRatesOfExchangeResponse{
		RatesOfExchange: result,
	}, err
}

func (s *storageReaderService) GetRateOfExchange(ctx context.Context, in *pb.GetRateOfExchangeRequest) (*pbm.RateOfExchange, error) {

	return pbm.DefaultReadRateOfExchange(ctx, &pbm.RateOfExchange{Id: in.GetId()}, storageReader.gormDB)
}

func (s *storageReaderService) ListLocations(ctx context.Context, in *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {

	result, err := pbm.DefaultListLocation(ctx, storageReader.gormDB)

	return &pb.ListLocationsResponse{
		Locations: result,
	}, err
}

func (s *storageReaderService) GetLocation(ctx context.Context, in *pb.GetLocationRequest) (*pbm.Location, error) {

	return pbm.DefaultReadLocation(ctx, &pbm.Location{Id: in.GetId()}, storageReader.gormDB)
}

func (s *storageReaderService) ListProducts(ctx context.Context, in *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {

	result, err := pbm.DefaultListProduct(ctx, storageReader.gormDB)

	return &pb.ListProductsResponse{
		Products: result,
	}, err
}

func (s *storageReaderService) GetProduct(ctx context.Context, in *pb.GetProductRequest) (*pbm.Product, error) {

	return pbm.DefaultReadProduct(ctx, &pbm.Product{Id: in.GetId()}, storageReader.gormDB)
}

func (s *storageReaderService) ListPrices(ctx context.Context, in *pb.ListPricesRequest) (*pb.ListPricesResponse, error) {

	result, err := pbm.DefaultListPrice(ctx, storageReader.gormDB)

	return &pb.ListPricesResponse{
		Prices: result,
	}, err
}

func (s *storageReaderService) GetPrice(ctx context.Context, in *pb.GetPriceRequest) (*pbm.Price, error) {

	return pbm.DefaultReadPrice(ctx, &pbm.Price{Id: in.GetId()}, storageReader.gormDB)
}

func (s *storageReaderService) ListPositions(ctx context.Context, in *pb.ListPositionsRequest) (*pb.ListPositionsResponse, error) {

	result, err := pbm.DefaultListPosition(ctx, storageReader.gormDB)

	return &pb.ListPositionsResponse{
		Positions: result,
	}, err
}

func (s *storageReaderService) GetPosition(ctx context.Context, in *pb.GetPositionRequest) (*pbm.Position, error) {

	return pbm.DefaultReadPosition(ctx, &pbm.Position{Id: in.GetId()}, storageReader.gormDB)
}

func (s *storageReaderService) ListEmployees(ctx context.Context, in *pb.ListEmployeesRequest) (*pb.ListEmployeesResponse, error) {

	result, err := pbm.DefaultListEmployee(ctx, storageReader.gormDB)

	return &pb.ListEmployeesResponse{
		Employees: result,
	}, err
}

func (s *storageReaderService) GetEmployee(ctx context.Context, in *pb.GetEmployeeRequest) (*pbm.Employee, error) {

	return pbm.DefaultReadEmployee(ctx, &pbm.Employee{Id: in.GetId()}, storageReader.gormDB)
}

func (s *storageReaderService) ListMethods(ctx context.Context, in *pb.ListMethodsRequest) (*pb.ListMethodsResponse, error) {

	result, err := pbm.DefaultListMethod(ctx, storageReader.gormDB)

	return &pb.ListMethodsResponse{
		Methods: result,
	}, err
}

func (s *storageReaderService) GetMethod(ctx context.Context, in *pb.GetMethodRequest) (*pbm.Method, error) {

	return pbm.DefaultReadMethod(ctx, &pbm.Method{Id: in.GetId()}, storageReader.gormDB)
}

func (s *storageReaderService) ListReceipts(ctx context.Context, in *pb.ListReceiptsRequest) (*pb.ListReceiptsResponse, error) {

	result, err := pbm.DefaultListReceipt(ctx, storageReader.gormDB)

	return &pb.ListReceiptsResponse{
		Receipts: result,
	}, err
}

func (s *storageReaderService) GetReceipt(ctx context.Context, in *pb.GetReceiptRequest) (*pbm.Receipt, error) {

	return pbm.DefaultReadReceipt(ctx, &pbm.Receipt{Id: in.GetId()}, storageReader.gormDB)
}

func (s *storageReaderService) ListPurchases(ctx context.Context, in *pb.ListPurchasesRequest) (*pb.ListPurchasesResponse, error) {

	result, err := pbm.DefaultListPurchase(ctx, storageReader.gormDB)

	return &pb.ListPurchasesResponse{
		Purchases: result,
	}, err
}

func (s *storageReaderService) GetPurchase(ctx context.Context, in *pb.GetPurchaseRequest) (*pbm.Purchase, error) {

	return pbm.DefaultReadPurchase(ctx, &pbm.Purchase{Id: in.GetId()}, storageReader.gormDB)
}
