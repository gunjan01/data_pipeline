package search

import (
	"context"

	"github.com/gunjan01/data_pipeline/source/config"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

// Es is a wrapper to elasticsearch client.
type Es struct {
	Client *elastic.Client
	ctx    context.Context
}

// New returns an instance of elasticsearch client.
func New() (*Es, error) {
	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200.
	client, err := elastic.NewClient(
		elastic.SetURL(config.ElasticSearchURL),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)
	if err != nil {
		logrus.Errorf("Connection to elastic search failed with %+v", err)
		return nil, err
	}

	return &Es{
		Client: client,
		ctx:    context.Background(),
	}, nil
}
