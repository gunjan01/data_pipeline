package config

var (
	// ElasticSearchURL is URL to elastic search.
	ElasticSearchURL string = "http://localhost:9200/"

	// ElasticsearchUsername is the username to connect to Elasticsearch.
	ElasticsearchUsername string

	// ElasticsearchPassword is the password to connect to Elasticsearch.
	ElasticsearchPassword string

	// SearchIndex is the elastic index that will be queried.
	SearchIndex string = "test"
)
