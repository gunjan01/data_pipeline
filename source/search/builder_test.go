package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getJSONFileData(t *testing.T, source interface{}, fixtureName string) (string, string) {
	data, err := json.Marshal(source)
	assert.Nil(t, err)

	file, err := ioutil.ReadFile(fmt.Sprintf("./fixtures/%s.golden", fixtureName))
	assert.Nil(t, err)

	return string(file), string(data)
}

func TestBuildDateRangeQuery(t *testing.T) {
	tests := map[string]struct {
		startDate   string
		endDate     string
		fixtureName string
	}{
		"Basic request with both upper and lower limits": {
			startDate:   "2020-06-01",
			endDate:     "2020-06-02",
			fixtureName: "TestDateRangeQuery_1",
		},
		"Basic request with only lower limit": {
			startDate:   "2020-06-01",
			fixtureName: "TestDateRangeQuery_2",
		},
		"Basic reauest with only upper limit": {
			endDate:     "2020-06-02",
			fixtureName: "TestDateRangeQuery_3",
		},
	}

	for testCase, test := range tests {
		t.Run(testCase, func(t *testing.T) {

			query := buildDateRangeQuery(test.startDate, test.endDate)
			if query != nil {
				src, err := query.Source()
				assert.Nil(t, err)

				jsonFile, jsonData := getJSONFileData(t, src, test.fixtureName)

				assert.JSONEq(t, jsonFile, jsonData)
			} else {
				assert.Nil(t, query)
			}
		})
	}
}

func TestBuildBreakdown(t *testing.T) {
	tests := map[string]struct {
		startDate   string
		endDate     string
		fixtureName string
	}{
		"Basic request for with upper and lower limits": {
			startDate:   "2020-06-01",
			endDate:     "2020-06-02",
			fixtureName: "TestQueryBreakdown_1",
		},
	}

	for testCase, test := range tests {
		t.Run(testCase, func(t *testing.T) {

			breakdown := buildBreakdown(test.startDate, test.endDate)
			if breakdown != nil {
				src, err := breakdown.Source()
				assert.Nil(t, err)

				jsonFile, jsonData := getJSONFileData(t, src, test.fixtureName)

				assert.JSONEq(t, jsonFile, jsonData)
			} else {
				assert.Nil(t, breakdown)
			}
		})
	}
}

func TestNewSearchSourceBuilder(t *testing.T) {
	tests := map[string]struct {
		startDate   string
		endDate     string
		fixtureName string
	}{
		"Basic test covering the happy path": {
			startDate:   "2020-06-01",
			endDate:     "2020-06-03",
			fixtureName: "TestSearchSourceBuilder_1",
		},
		"Basic test with only lower limit of date range set": {
			startDate:   "2020-06-01",
			fixtureName: "TestSearchSourceBuilder_2",
		},
		"Basic test with only upper limit of date range set": {
			endDate:     "2020-06-03",
			fixtureName: "TestSearchSourceBuilder_3",
		},
	}

	for testCase, test := range tests {
		t.Run(testCase, func(t *testing.T) {

			searchSource := NewSearchSourceBuilder(test.startDate, test.endDate)
			if searchSource != nil {
				src, err := searchSource.Source()
				assert.Nil(t, err)

				jsonFile, jsonData := getJSONFileData(t, src, test.fixtureName)

				assert.JSONEq(t, jsonFile, jsonData)
			} else {
				assert.Nil(t, searchSource)
			}
		})
	}
}
