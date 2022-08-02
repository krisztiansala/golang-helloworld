.DEFAULT_GOAL = all
all: test vet fmt build

test:
	go test ./...
vet:
	go vet ./...
fmt:
	go fmt ./...
build:
	go build -o bin/ .
dev:
	nodemon --exec go run main.go --signal SIGTERM
