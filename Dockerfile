FROM golang:1.18-buster AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY ./ ./
RUN make build

FROM gcr.io/distroless/static-debian11
WORKDIR /
EXPOSE 8080
USER nonroot:nonroot
COPY --from=build --chown=nonroot:nonroot /app/bin/golang-helloworld /golang-hello
ENTRYPOINT ["/golang-hello"]
