package search

import elastic "github.com/olivere/elastic"

// NewSearchSourceBuilder builds the elastic searchsource.
func NewSearchSourceBuilder(startDate string, endDate string) *elastic.SearchSource {
	searchSource := elastic.NewSearchSource().Size(999)
	searchSource.Aggregation("queries", buildBreakdown(startDate, endDate))

	return searchSource
}

// buildBreakdown builds the nested aggregation. We aggregate the results on the query and apply
// a sum aggregation after filtering it for a specific date range.
func buildBreakdown(startDate string, endDate string) *elastic.FilterAggregation {
	// Apply a sum operation on each count field.
	countSubAggregation := elastic.NewSumAggregation().Field("count")

	queryAggregation := elastic.NewTermsAggregation().
		Field("query.keyword").
		Order("_count", false).
		Size(1000)
	queryAggregation.SubAggregation("bucket_sum", countSubAggregation)

	filtersAggregation := elastic.NewFilterAggregation().
		Filter(
			buildDateRangeQuery(startDate, endDate),
		).SubAggregation("query_breakdown", queryAggregation)

	return filtersAggregation
}

// buildDateRangeQuery builds a range query and extracts results based on a
// specific date range.
func buildDateRangeQuery(startDate string, endDate string) *elastic.BoolQuery {
	if len(startDate) == 0 {
		return elastic.NewBoolQuery().Must(
			elastic.NewRangeQuery("date").Lte(endDate),
		)
	}

	if len(endDate) == 0 {
		return elastic.NewBoolQuery().Must(
			elastic.NewRangeQuery("date").Gte(startDate),
		)
	}

	return elastic.NewBoolQuery().Must(
		elastic.NewRangeQuery("date").Gte(startDate).Lte(endDate),
	)
}
