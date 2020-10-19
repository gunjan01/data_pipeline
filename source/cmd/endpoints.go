package main

import (
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/gunjan01/data_pipeline/source/httpapi"
)

func endpoints(api *httpapi.HTTPAPI) http.Handler {
	mux := bone.New()

	mux.Get(
		"/data",
		http.HandlerFunc(api.GetCollatedData),
	)

	return mux
}
