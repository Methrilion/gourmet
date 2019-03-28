build:
		docker build -t gourmet-db ./db

unfail:
		go get -u github.com/methrilion/gourmet

run:
		docker-compose up

down:
		docker-compose down

clean:
		docker rm gourmet-db

re:
		make down
		make build
		make run

# ifndef $(GOPATH)
#     GOPATH=$(shell go env GOPATH)
#     export GOPATH
# endif
