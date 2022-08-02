.DEFAULT_GOAL = all
all: test vet fmt build

test:
	@go test ./...
vet:
	go vet ./...
fmt:
	go fmt ./...
build:
	go build -o bin/ .
