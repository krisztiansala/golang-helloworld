.DEFAULT_GOAL = all
all: test vet fmt build

test:
	go test ./...
vet:
	go vet ./...
fmt:
	go fmt ./...
build:
	go build -ldflags="-X main.GitHash=`git rev-parse HEAD` -X main.GitProject=`git remote get-url origin | xargs basename -s .git`" -o bin/ ./...
dev:
	cd cmd/golang-helloworld && nodemon --exec go run . --port=4000 --signal SIGTERM
run:
	cd cmd/golang-helloworld && go run .
deploy:
	kubectl apply -f k8s
helm_deploy:
	helm upgrade --install golang-helloworld --values helm/values.yaml helm
helm_uninstall:
	helm uninstall golang-helloworld
delete:
	kubectl delete -f k8s
portforward:
	kubectl port-forward service/golang-helloworld-service 8080:8080
tf_local_apply:
	terraform -chdir=terraform/local apply -auto-approve
tf_gcp_apply:
	terraform -chdir=terraform/gcp apply -auto-approve
tf_local_destroy:
	terraform -chdir=terraform/local destroy -auto-approve
tf_gcp_destroy:
	terraform -chdir=terraform/gcp destroy -auto-approve
