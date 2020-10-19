package main

import (
	"fmt"
	"net/http"

	"github.com/gunjan01/data_pipeline/source/config"
	"github.com/gunjan01/data_pipeline/source/httpapi"
	"github.com/sirupsen/logrus"
)

func main() {
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

// Dont repetadely load csv. Use action => 'update'
