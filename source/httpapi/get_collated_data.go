package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/gunjan01/data_pipeline/source/config"
	"github.com/gunjan01/data_pipeline/source/search"
	"github.com/sirupsen/logrus"
)

// HTTPAPI is the base struct for the handlers. It's what
// the mux router will use when instantiating the HTTP handlers.
type HTTPAPI struct {
}

// GetCollatedData returns the squashed data for a particular date range.
func (h *HTTPAPI) GetCollatedData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if len(startDate) == 0 && len(endDate) == 0 {
		logrus.Errorf("You must specify a date-range for the endpoint to work")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	client, err := search.New()
	if err != nil {
		panic(err)
	}

	searchSource := search.NewSearchSourceBuilder(startDate, endDate)
	response, err := client.ParseResults(config.SearchIndex, searchSource)
	if err != nil {
		panic(err)
	}

	logrus.Infof("%+v", response)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logrus.Errorf("Failed to construct response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
