package search

import (
	"errors"

	elastic "github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

type teststruct struct {
	query string
	count float64
}

type ResponseStruct struct {
	test []teststruct
}

func (c *Es) searchResultHasErr(searchErr *elastic.ErrorDetails) error {
	if searchErr == nil {
		return nil
	}

	return errors.New(searchErr.Reason)
}

// ParseResults will make the call to elastic and construct the response.
func (c *Es) ParseResults(index string, searchSource *elastic.SearchSource) (ResponseStruct, error) {
	logrus.Infof(index)
	response := ResponseStruct{}
	source, err := searchSource.Source()
	if err != nil {
		return response, err
	}

	search := c.Client.Search(index).Source(source)
	result, err := search.Do(c.ctx)
	if err != nil {
		return response, err
	}

	grouping, err := extractResult(result)
	if err != nil {
		return response, err
	}

	err = c.searchResultHasErr(result.Error)
	if err != nil {
		return response, err
	}

	response.test = grouping
	return response, nil
}

// extarct Result extracts the results from the ES results.
func extractResult(result *elastic.SearchResult) ([]teststruct, error) {
	response := []teststruct{}
	var count float64
	breakdown, ok := result.Aggregations.Filter("queries")
	if ok {
		if queryBreakdown, ok := breakdown.Aggregations.Terms("query_breakdown"); ok {
			buckets := queryBreakdown.Buckets

			for _, bucket := range buckets {
				// Extract the sum Aggregation
				sumAggregation, ok := bucket.Aggregations.Sum("bucket_sum")
				if ok {
					value := *sumAggregation.Value
					if value < 50 {
						count = count + value
						continue
					}

					responseResult := teststruct{
						query: bucket.Key.(string),
						count: value,
					}
					response = append(response, responseResult)
				}
			}

			response = append(response, teststruct{
				query: "Query < 50",
				count: count,
			})
		}
	}

	return response, nil
}
