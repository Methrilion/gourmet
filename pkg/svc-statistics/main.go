package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	pbreader "github.com/methrilion/gourmet/proto/reader"
	pb "github.com/methrilion/gourmet/proto/statistics"
	"google.golang.org/grpc"
)

type statisticsService struct {
	reader pbreader.StorageReaderServiceClient
}

var (
	statistics            statisticsService
	storageReaderEndpoint = os.Getenv("STORAGE_READER_SERVICE_FULLADDRESS")
)

func main() {

	time.Sleep(5 * time.Second) ///////////////// TODO : CHANGE TO healthcheck IN docker-compose

	opts := []grpc.DialOption{grpc.WithInsecure()}

	log.Println("Connecting to services...")
	conn, err := grpc.Dial(storageReaderEndpoint, opts...)
	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}
	defer conn.Close()
	log.Println("Connection successful")

	c := pbreader.NewStorageReaderServiceClient(conn)
	statistics.reader = c

	/////////////////

	/////////////////

	runTestFuncs()

	/////////////////

	/////////////////

	lis, err := net.Listen("tcp", os.Getenv("STATISTICS_SERVICE_ADDRESS"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	pb.RegisterStatisticsServiceServer(s, &statistics)

	/////////////////

	/////////////////

	log.Println("Now listening on", os.Getenv("STATISTICS_SERVICE_ADDRESS"))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

func runTestFuncs() {
	ctxBack := context.Background()

	////////////////////// 1 PART
	response1, err := statistics.reader.GetCountPurchasesByYear(ctxBack, &pbreader.GetCountPurchasesByYearRequest{Year: 2016})
	if err != nil {
		log.Fatalf("Error when calling GetCountPurchasesByYear: %s", err)
	}
	log.Println("Response from server:", response1.Counts)

	////////////////////// 2 PART
	start, _ := ptypes.TimestampProto(time.Date(2016, 7, 1, 0, 0, 0, 0, time.UTC))
	end, _ := ptypes.TimestampProto(time.Date(2016, 8, 1, 0, 0, 0, -1, time.UTC))
	response2, err := statistics.reader.GetRevenuePurchasesByDatesByPrice(ctxBack,
		&pbreader.GetRevenuePurchasesByDatesByPriceRequest{
			PriceId: 3,
			Start:   start,
			End:     end,
		})
	if err != nil {
		log.Fatalf("Error when calling GetRevenuePurchasesByDatesByPrice: %s", err)
	}
	log.Println("Response from server:", response2.Revenue)

	////////////////////// 3 PART
	response3, err := statistics.reader.GetRevenuePurchasesByDatesByProduct(ctxBack,
		&pbreader.GetRevenuePurchasesByDatesByProductRequest{
			ProductId: 3,
			Start:     start,
			End:       end,
		})
	if err != nil {
		log.Fatalf("Error when calling GetRevenuePurchasesByDatesByProduct: %s", err)
	}
	log.Println("Response from server:", response3.Revenue)
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

func (s *statisticsService) GetPricesStatisticsByYear(ctx context.Context, in *pb.GetPricesStatisticsByYearRequest) (*pb.GetPricesStatisticsByYearResponse, error) {

	type pricesStat struct {
		id       uint32
		revenues []float32
		sum      float32
	}

	ctxBack := context.Background()
	response, err := statistics.reader.ListPrices(ctxBack, &pbreader.ListPricesRequest{})
	if err != nil {
		return nil, err
	}

	starts, ends := getMonths(int(in.GetYear()))

	var pricesStats []pricesStat

	for _, a := range response.Prices {
		id := a.GetId()
		tmp := pricesStat{id: id}

		for j := 0; j < 12; j++ {
			start, _ := ptypes.TimestampProto(starts[j])
			end, _ := ptypes.TimestampProto(ends[j])

			response2, err := statistics.reader.GetRevenuePurchasesByDatesByPrice(ctxBack,
				&pbreader.GetRevenuePurchasesByDatesByPriceRequest{
					PriceId: id,
					Start:   start,
					End:     end,
				})
			if err != nil {
				return nil, err
			}
			log.Println(response2.Revenue)

			tmp.revenues = append(tmp.revenues, response2.Revenue)
			tmp.sum += response2.Revenue
		}
		pricesStats = append(pricesStats, tmp)
	}

	var results [12]float32
	for i := 0; i < 12; i++ {
		var sum float32
		for _, a := range pricesStats {
			sum += a.revenues[i]
		}
		results[i] = sum
	}

	log.Println(pricesStats)

	var pbresult []*pb.PricesStat
	for _, a := range pricesStats {
		pbresult = append(pbresult,
			&pb.PricesStat{
				Id:       a.id,
				Revenues: a.revenues,
				Sum:      a.sum,
			})
	}
	return &pb.GetPricesStatisticsByYearResponse{
		PricesStat: pbresult,
		Results:    results[:],
	}, nil
}

func (s *statisticsService) GetProductsStatisticsByYear(ctx context.Context, in *pb.GetProductsStatisticsByYearRequest) (*pb.GetProductsStatisticsByYearResponse, error) {

	type productsStat struct {
		id       uint32
		revenues []float32
		sum      float32
	}

	ctxBack := context.Background()
	pbListProducts, err := statistics.reader.ListProducts(ctxBack, &pbreader.ListProductsRequest{})
	if err != nil {
		return nil, err
	}

	starts, ends := getMonths(int(in.GetYear()))

	var productsStats []*pb.ProductsStat

	for _, a := range pbListProducts.Products {
		id := a.GetId()
		tmp := &pb.ProductsStat{Id: id}

		for j := 0; j < 12; j++ {
			start, _ := ptypes.TimestampProto(starts[j])
			end, _ := ptypes.TimestampProto(ends[j])

			res, err := statistics.reader.GetRevenuePurchasesByDatesByProduct(ctxBack,
				&pbreader.GetRevenuePurchasesByDatesByProductRequest{
					ProductId: id,
					Start:     start,
					End:       end,
				})
			if err != nil {
				return nil, err
			}

			tmp.Revenues = append(tmp.Revenues, res.Revenue)
			tmp.Sum += res.Revenue
		}
		productsStats = append(productsStats, tmp)
	}

	var results [12]float32
	for i := 0; i < 12; i++ {
		var revenueByMonth float32
		for _, a := range productsStats {
			revenueByMonth += a.Revenues[i]
		}
		results[i] = revenueByMonth
	}

	log.Println(productsStats)

	return &pb.GetProductsStatisticsByYearResponse{
		ProductsStat: productsStats,
		Results:      results[:],
	}, nil
}
