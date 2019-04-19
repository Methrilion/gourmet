package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/methrilion/gourmet/proto/statistics"
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

	c := pbstatistics.NewStatisticsServiceClient(conn)
	analytics.statistics = c

	ctx := context.Background()
	res, _ := analytics.statistics.GetPricesStatisticsByYear(ctx, &pb.GetPricesStatisticsByYearRequest{Year: 2016})
	log.Println(res)
	res2, _ := analytics.statistics.GetProductsStatisticsByYear(ctx, &pb.GetProductsStatisticsByYearRequest{Year: 2016})
	log.Println(res2)

	/////////////////

	/////////////////

	// runTestFuncs()
	// statistics.GetPricesStatisticsByYear(nil, &pb.GetPricesStatisticsByYearRequest{Year: 2016})
	// statistics.GetProductsStatisticsByYear(nil, &pb.GetProductsStatisticsByYearRequest{Year: 2016})

	/////////////////

	/////////////////

	// lis, err := net.Listen("tcp", os.Getenv("STATISTICS_SERVICE_ADDRESS"))
	// if err != nil {
	// 	log.Fatalf("Failed to listen: %v", err)
	// }
	// defer lis.Close()

	// s := grpc.NewServer()
	// pb.RegisterStatisticsServiceServer(s, &statistics)

	// /////////////////

	// /////////////////

	// log.Println("Now listening on", os.Getenv("STATISTICS_SERVICE_ADDRESS"))

	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("Failed to serve: %v\n", err)
	// }
}
