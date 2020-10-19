package main

import (
	"fmt"

	"github.com/gunjan01/data_pipeline/source/config"
	"github.com/gunjan01/data_pipeline/source/search"
)

func main() {
	client, err := search.New()
	if err != nil {
		panic(err)
	}

	searchSource := search.NewSearchSourceBuilder("2020-06-01", "2020-06-04")
	response, err := client.ParseResults(config.SearchIndex, searchSource)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

// Dont repetadely load csv. Use action => 'update'
