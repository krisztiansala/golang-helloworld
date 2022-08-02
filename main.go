package main

import (
	"flag"
	"fmt"
	"os"

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

func main() {
	port := flag.Int("port", 8080, "The port the server should run on. Default value is 8080.")
	flag.Parse()
	fmt.Println(*port)
	log.Infof("Starting application on port %d", *port)
}
