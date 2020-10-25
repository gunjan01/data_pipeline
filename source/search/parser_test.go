package search

import (
	"encoding/json"
	"testing"

	elastic "github.com/olivere/elastic"
	"github.com/stretchr/testify/assert"
)

func TestExtractResult(t *testing.T) {
	tests := map[string]struct {
		result           *elastic.SearchResult
		expectedResponse []CollateData
		err              error
		BreakdownLen     int
	}{
		"Basic test": {
			result: &elastic.SearchResult{
				TookInMillis: 8,
				Aggregations: func() elastic.Aggregations {
					raw := json.RawMessage(`{"doc_count":8,"query_breakdown":{"doc_count_error_upper_bound":0,"sum_other_doc_count":0,"buckets":[{"key":"best web hosting","doc_count":2,"bucket_sum":{"value":4000.0}},{"key":"cheap web hosting","doc_count":2,"bucket_sum":{"value":8000.0}},{"key":"test","doc_count":2,"bucket_sum":{"value":2.0}},{"key":"test1","doc_count":2,"bucket_sum":{"value":1.0}}]}}`)
					return map[string]*json.RawMessage{
						"queries": &raw,
					}
				}(),
			},
			expectedResponse: []CollateData{
				CollateData{
					Query: "best web hosting",
					Count: 4000.0,
				},
				CollateData{
					Query: "cheap web hosting",
					Count: 8000.0,
				},
				CollateData{
					Query: "Query < 50",
					Count: 3.0,
				},
			},
			BreakdownLen: 3,
		},
		"no breakdown present": {
			result: &elastic.SearchResult{
				TookInMillis: 100,
				Aggregations: func() elastic.Aggregations {
					raw := json.RawMessage(`{"doc_count":8,"query_breakdown":{"doc_count_error_upper_bound":0,"sum_other_doc_count":0,"buckets":[]}}`)
					return map[string]*json.RawMessage{
						"queries": &raw,
					}
				}(),
			},
			expectedResponse: []CollateData{},
			BreakdownLen:     0,
		},
	}

	for testCase, test := range tests {
		t.Run(testCase, func(t *testing.T) {
			response, err := ExtractResult(test.result)

			if err != nil {
				t.Errorf("Test failed with err: %v", err)
			}

			assert.Equal(t, response, test.expectedResponse)
			assert.Equal(t, len(response), test.BreakdownLen)
		})
	}
}
