package httpapi

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gunjan01/data_pipeline/source/search"
	"github.com/stretchr/testify/assert"
)

func getDataRequest(startDate string, endDate string) *http.Request {
	req, _ := http.NewRequest("GET", "/data", nil)
	q := url.Values{}
	q.Add("start_date", startDate)
	q.Add("end_date", endDate)
	req.URL.RawQuery = q.Encode()

	return req
}
func TestCreateContact(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tests := map[string]struct {
		startDate        string
		endDate          string
		expectedResponse []search.CollateData
		expectedCode     int
	}{
		"Bad request": {
			expectedCode: http.StatusBadRequest,
		},
		"request with both start and end dates": {
			startDate: "2020-06-01",
			endDate:   "2020-06-02",
			expectedResponse: []search.CollateData{
				search.CollateData{
					Query: "best web hosting",
					Count: 4000.0999755859375,
				},
				search.CollateData{
					Query: "cheap web hosting",
					Count: 8000,
				},
				search.CollateData{
					Query: "Query < 50",
					Count: 3,
				},
			},
			expectedCode: http.StatusOK,
		},
		"request with only start date": {
			startDate: "2020-06-30",
			expectedResponse: []search.CollateData{
				search.CollateData{
					Query: "best web hosting",
					Count: 2000,
				},
				search.CollateData{
					Query: "cheap web hosting",
					Count: 4000,
				},
				search.CollateData{
					Query: "Query < 50",
					Count: 1.5,
				},
			},
			expectedCode: http.StatusOK,
		},
		"request with only end date": {
			startDate: "2020-06-01",
			expectedResponse: []search.CollateData{
				search.CollateData{
					Query: "2000.0999755859375",
					Count: 1,
				},
				search.CollateData{
					Query: "4000",
					Count: 1,
				},
				search.CollateData{
					Query: "Query < 50",
					Count: 1.5,
				},
			},
			expectedCode: http.StatusOK,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {

			api := HTTPAPI{}

			w := httptest.NewRecorder()

			handler := http.HandlerFunc(api.GetCollatedData)
			handler.ServeHTTP(w, getDataRequest(test.startDate, test.endDate))
			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}
