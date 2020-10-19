package search

/*
func TestExtractResult(t *testing.T) {
	tests := map[string]struct {
		result       *elastic.SearchResult
		response     []teststruct
		err          error
		BreakdownLen int
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
			BreakdownLen: 2,
		},
		/*	"no breakdown test": {
			result: &elastic.SearchResult{
				TookInMillis: 100,
				Aggregations: func() elastic.Aggregations {
					cat := json.RawMessage(`{"doc_count_error_upper_bound":0,"sum_other_doc_count":0,"buckets":[]}`)
					return map[string]*json.RawMessage{
						"l1_locations": &cat,
						"l2_locations": &cat,
					}
				}(),
			},
			response: &searchV2.GetLocationBreakdownResponse{
				Breakdown: make(map[*config.Location]int32),
			},
			BreakdownLen: 0,
		},
	}

	for testCase, test := range tests {
		t.Run(testCase, func(t *testing.T) {
			response, err := ExtractResult(test.result)
			//assert.Len(t, parser.response.Breakdown, test.BreakdownLen)
		})
	}
}
*/
