package search

import (
	"github.com/olivere/elastic"
)

// NewSearchSourceBuilder builds the elastic searchsource.
func NewSearchSourceBuilder(startDate string, endDate string) *elastic.SearchSource {
	searchSource := elastic.NewSearchSource().Size(0)
	searchSource.Aggregation("queries", buildBreakdown(startDate, endDate))

	return searchSource
}

func buildBreakdown(startDate string, endDate string) *elastic.FilterAggregation {
	countSubAggregation := elastic.NewSumAggregation().Field("count")

	queryAggregation := elastic.NewTermsAggregation().
		Field("query").
		Order("_count", false).
		Size(1000)
	queryAggregation.SubAggregation("bucket_sum", countSubAggregation)

	filtersAggregation := elastic.NewFilterAggregation().
		Filter(
			buildDateRangeQuery(startDate, endDate),
		).SubAggregation("query_breakdown", queryAggregation)

	return filtersAggregation
}

func buildDateRangeQuery(startDate string, endDate string) *elastic.BoolQuery {
	return elastic.NewBoolQuery().Must(
		elastic.NewRangeQuery("date").Gte(startDate).Lte(endDate),
	)
}
