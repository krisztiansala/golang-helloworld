FROM golang:1.18-buster AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY ./ ./
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /golang-hello .

FROM gcr.io/distroless/static-debian11
WORKDIR /
EXPOSE 8080
USER nonroot:nonroot
COPY --from=build --chown=nonroot:nonroot /golang-hello /golang-hello
ENTRYPOINT ["/golang-hello"]
