package main

import (
	"context"
	"log"
	"net"
	"os"
	"sort"
	"time"

	pb "github.com/methrilion/gourmet/proto/analytics"
	pbstatistics "github.com/methrilion/gourmet/proto/statistics"
	"google.golang.org/grpc"
)

type analyticsService struct {
	statistics pbstatistics.StatisticsServiceClient
}

var (
	analytics          analyticsService
	statisticsEndpoint = os.Getenv("STATISTICS_SERVICE_FULLADDRESS")
)

func main() {

	time.Sleep(10 * time.Second) ///////////////// TODO : CHANGE TO healthcheck IN docker-compose

	opts := []grpc.DialOption{grpc.WithInsecure()}

	log.Println("Connecting to services...")
	conn, err := grpc.Dial(statisticsEndpoint, opts...)
	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}
	defer conn.Close()
	log.Println("Connection successful")

	analytics.statistics = pbstatistics.NewStatisticsServiceClient(conn)

	///////////////// TEST PART ////////////////////////////////////////////

	ctx := context.Background()
	res, _ := analytics.GetABCByPrices(ctx, &pb.GetABCByPricesRequest{Year: 2016})
	log.Println(res)

	///////////////////////////////////////////////////////////////////////

	lis, err := net.Listen("tcp", os.Getenv("ANALYTICS_SERVICE_ADDRESS"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	pb.RegisterAnalyticsServiceServer(s, &analytics)

	log.Println("Now listening on", os.Getenv("ANALYTICS_SERVICE_ADDRESS"))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

func getStringABC(state int) string {
	if state == 0 {
		return "A"
	}
	if state == 1 {
		return "B"
	}
	return "C"
}

func (s *analyticsService) GetABCByPrices(ctx context.Context, in *pb.GetABCByPricesRequest) (*pb.GetABCByPricesResponse, error) {
	ctxBack := context.Background()
	res, _ := analytics.statistics.GetPricesStatisticsByYear(ctxBack, &pbstatistics.GetPricesStatisticsByYearRequest{Year: in.GetYear()})

	pricesStats := res.PricesStat
	length := len(pricesStats)

	sort.Slice(pricesStats, func(i, j int) bool { return pricesStats[i].Sum > pricesStats[j].Sum })

	parts := []int{
		length / 3,
		2 * (length / 3),
		length,
	}
	result := []*pb.PriceABC{}

	var state int

	for i := 0; i < length; i++ {

		for j := 0; j < 3; j++ {
			if i < parts[j] {
				state = j
				break
			}
		}
		result = append(result, &pb.PriceABC{Id: pricesStats[i].Id, ABC: getStringABC(state)})
	}

	return &pb.GetABCByPricesResponse{
		PriceABC: result,
	}, nil
}
