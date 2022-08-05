package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	"github.com/krisztiansala/golang-helloworld/util"
	log "github.com/sirupsen/logrus"
)

var GitProject, GitHash string

type FlagConfig struct {
	port int
	args []string
}

func parseFlags(progname string, args []string) (config *FlagConfig, output string, err error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)

	var conf FlagConfig
	flags.IntVar(&conf.port, "port", 0, "The port the server should run on. Default value is 8080.")

	err = flags.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}
	conf.args = flags.Args()
	return &conf, buf.String(), nil
}
func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	flagConf, flagOutput, err := parseFlags(os.Args[0], os.Args[1:])

	// If the -h or --help flags are passed, it will print out the usage message
	if err == flag.ErrHelp {
		fmt.Println(flagOutput)
		os.Exit(2)
	} else if err != nil {
		fmt.Println("Error while parsing flags: ", err)
		fmt.Println("Output:\n", flagOutput)
		os.Exit(1)
	}

	portEnv, err := util.GetenvIntDefault("PORT", 0)
	if err != nil {
		fmt.Println("Error while parsing environment variable: ", err)
	}
	port := util.GetPort(flagConf.port, portEnv)

	environment := util.GetenvDefault("ENV", "production")
	listenAddress := "0.0.0.0"
	if environment == "development" {
		listenAddress = "127.0.0.1"
	}

	if err != nil {
		log.Fatal("Error occurred parsing port number from environmental variable")
	}
	log.Infof("Starting application on the %s environment, address %s:%d", environment, listenAddress, port)

	http.HandleFunc("/helloworld", helloHandler)
	http.HandleFunc("/versionz", versionHandler)
	http.HandleFunc("/", RootHandler)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", listenAddress, port),
		Handler: logRequest(http.DefaultServeMux),
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Error initializing server: ", err)
		}
	}()

	<-ctx.Done()
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	log.Info("Server shutting down gracefully...")

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forcefully shut down: ", err)
	}

}
