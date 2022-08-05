.DEFAULT_GOAL = all
all: test vet fmt build

test:
	go test ./...
vet:
	go vet ./...
fmt:
	go fmt ./...
build:
	go build -ldflags="-X main.GitHash=`git rev-parse HEAD` -X main.GitProject=`git remote get-url origin | xargs basename -s .git`" -o bin/ .
dev:
	nodemon --exec go run . --port=4000 --signal SIGTERM
