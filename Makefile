build:
		GOOS=linux GOARCH=amd64 go build -o ./pkg/svc-storage-writer/svc-storage-writer -i ./pkg/svc-storage-writer/*.go
		docker build -t svc-storage-writer ./pkg/svc-storage-writer
		docker build -t gourmet-db ./db
		rm ./pkg/svc-storage-writer/svc-storage-writer

unfail:
		go get -u github.com/methrilion/gourmet

run:
		docker-compose up

down:
		docker-compose down

clean:
		docker rm svc-storage-writer gourmet-db

re:
		make down
		make build
		make run

# ifndef $(GOPATH)
#     GOPATH=$(shell go env GOPATH)
#     export GOPATH
# endif
