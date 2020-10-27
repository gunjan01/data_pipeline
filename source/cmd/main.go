package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gunjan01/data_pipeline/source/config"
	"github.com/gunjan01/data_pipeline/source/httpapi"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing parameter, provide file name!")
		return
	}

	// Manually load the CSV into the elastic index.
	// Ideally logstash should be employed to do this.
	LoadCSV(os.Args[1])

	// Start the http server on port 9090.
	address := fmt.Sprintf(":%d", config.Port)

	logrus.WithFields(logrus.Fields{
		"address": address,
	}).Info("Starting http service")

	api := httpapi.HTTPAPI{}

	server := http.Server{}
	server.Addr = address
	server.Handler = endpoints(&api)

	err := server.ListenAndServe()
	if err != nil {
		logrus.WithError(err).Panic("Failed to start the http server")
	}
}
