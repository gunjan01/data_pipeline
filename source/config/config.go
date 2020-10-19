package config

var (
	// Port the server is running on.
	Port int = 9090

	// ElasticSearchURL is URL to elastic search.
	ElasticSearchURL string = "http://localhost:9200/"

	// ElasticsearchUsername is the username to connect to Elasticsearch.
	ElasticsearchUsername string

	// ElasticsearchPassword is the password to connect to Elasticsearch.
	ElasticsearchPassword string

	// SearchIndex is the elastic index that will be queried.
	SearchIndex string = "test"

	// DefaultSize is the default size used in aggregations.
	DefaultSize int = 9999
)
