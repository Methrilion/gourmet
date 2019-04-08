package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "github.com/methrilion/gourmet/proto/writer"
)

var (
	storageWriterEndpoint = flag.String("svc-storage-writer_endpoint", "svc-storage-writer:9091", "endpoint of Storage Writer Service")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// err := gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)

	log.Println("Connecting to services...")
	if err := gw.RegisterStorageWriterServiceHandlerFromEndpoint(ctx, mux, *storageWriterEndpoint, opts); err != nil {
		return err
	}
	log.Println("Connection successful")

	log.Println("Now listening on", os.Getenv("INTEGRATOR_SERVICE_ADDRESS"))
	return http.ListenAndServe(os.Getenv("INTEGRATOR_SERVICE_ADDRESS"), mux)
}

func main() {

	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
