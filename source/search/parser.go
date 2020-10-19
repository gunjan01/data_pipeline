package search

import (
	"errors"

	elastic "github.com/olivere/elastic"
)

// CollateData contains the collated data
// for a particular data-range.
type CollateData struct {
	Query string  `json:"query"`
	Count float64 `json:"count"`
}

func (c *Es) searchResultHasErr(searchErr *elastic.ErrorDetails) error {
	if searchErr == nil {
		return nil
	}

	return errors.New(searchErr.Reason)
}

// ParseResults will make the actual call to elastic, parse the result and construct the response.
func (c *Es) ParseResults(index string, searchSource *elastic.SearchSource) ([]CollateData, error) {
	response := []CollateData{}
	source, err := searchSource.Source()
	if err != nil {
		return response, err
	}

	search := c.Client.Search(index).Source(source)
	result, err := search.Do(c.ctx)
	if err != nil {
		return response, err
	}

	collatedData, err := ExtractResult(result)
	if err != nil {
		return response, err
	}

	err = c.searchResultHasErr(result.Error)
	if err != nil {
		return response, err
	}

	return collatedData, nil
}

// ExtractResult extracts the results from the ES results.
func ExtractResult(result *elastic.SearchResult) ([]CollateData, error) {
	response := []CollateData{}
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

					responseResult := CollateData{
						Query: bucket.Key.(string),
						Count: value,
					}
					response = append(response, responseResult)
				}
			}

			response = append(response, CollateData{
				Query: "Query < 50",
				Count: count,
			})
		}
	}

	return response, nil
}
