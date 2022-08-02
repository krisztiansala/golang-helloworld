package main

import (
	"flag"
	"fmt"
	"os"

	"net/http"

	"github.com/krisztiansala/golang-helloworld/util"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Stranger")
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	portEnv, err := util.GetenvIntDefault("PORT", 8080)
	environment := util.GetenvDefault("ENV", "production")
	listenAddress := "0.0.0.0"
	if environment == "development" {
		listenAddress = "127.0.0.1"
	}

	if err != nil {
		log.Fatal("Error occurred parsing port number from environmental variable")
	}
	port := flag.Int("port", portEnv, "The port the server should run on. Default value is 8080.")
	flag.Parse()
	log.Infof("Starting application on the %s environment, port %d", listenAddress, *port)

	http.HandleFunc("/", rootHandler)

	err = http.ListenAndServe(fmt.Sprintf("%s:%d", listenAddress, *port), logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}
