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
deploy:
	kubectl apply -f k8s
helm_deploy:
	helm upgrade --install golang-helloworld --values helm/values.yaml helm
delete:
	kubectl delete -f k8s
portforward:
	kubectl port-forward service/golang-helloworld-service 8080:8080
