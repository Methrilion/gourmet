build:
		GOOS=linux GOARCH=amd64 go build -o ./pkg/svc-storage-writer/svc-storage-writer -i ./pkg/svc-storage-writer/*.go
		GOOS=linux GOARCH=amd64 go build -o ./pkg/svc-storage-reader/svc-storage-reader -i ./pkg/svc-storage-reader/*.go
		GOOS=linux GOARCH=amd64 go build -o ./pkg/svc-integrator/svc-integrator -i ./pkg/svc-integrator/*.go
		GOOS=linux GOARCH=amd64 go build -o ./pkg/svc-statistics/svc-statistics -i ./pkg/svc-statistics/*.go
		docker build -t gourmet-db ./db
		docker build -t svc-storage-writer ./pkg/svc-storage-writer
		docker build -t svc-storage-reader ./pkg/svc-storage-reader
		docker build -t svc-integrator ./pkg/svc-integrator
		docker build -t svc-statistics ./pkg/svc-statistics
		rm ./pkg/svc-storage-writer/svc-storage-writer
		rm ./pkg/svc-storage-reader/svc-storage-reader
		rm ./pkg/svc-integrator/svc-integrator
		rm ./pkg/svc-statistics/svc-statistics

unfail:
		go get -u github.com/methrilion/gourmet

run:
		docker-compose up

down:
		docker-compose down

clean:
		docker rm svc-storage-writer svc-storage-reader svc-integrator svc-statistics gourmet-db

re:
		make down
		make build
		make run

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

protobuf:
		protoc -I. \
			-I$(GOPATH)/src \
			--go_out=. \
			--gorm_out="engine=postgres:." \
			proto/model/model.proto
		protoc -I. \
			-I$(GOPATH)/src \
			-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--go_out=plugins=grpc:. \
			proto/writer/writer.proto
		protoc -I. \
			-I/usr/local/include \
			-I$(GOPATH)/src \
			-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--grpc-gateway_out=logtostderr=true:. \
			proto/writer/writer.proto
		protoc -I. \
			-I$(GOPATH)/src \
			-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--go_out=plugins=grpc:. \
			proto/reader/reader.proto
		protoc -I. \
			-I$(GOPATH)/src \
			--go_out=plugins=grpc:. \
			proto/statistics/statistics.proto

