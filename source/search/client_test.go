package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	elasticCient, err := New()
	if err != nil {
		t.Errorf("Test failed with err: %v", err)
	}

	assert.NotNil(t, elasticCient.ctx)
	assert.NotNil(t, elasticCient.Client)
}
