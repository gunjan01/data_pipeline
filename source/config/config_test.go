package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := map[string]struct {
		expectedValue interface{}
		variable      interface{}
	}{
		"tests whether ElasticSearchURL is set": {
			expectedValue: "http://localhost:9200/",
			variable:      ElasticSearchURL,
		},
		"tests whether SearchIndex is set": {
			expectedValue: "test",
			variable:      SearchIndex,
		},
		"tests whether DefaultSize is set": {
			expectedValue: 9999,
			variable:      DefaultSize,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expectedValue, test.variable)
		})
	}
}
