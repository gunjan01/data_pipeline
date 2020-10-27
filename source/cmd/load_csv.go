package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/gunjan01/data_pipeline/source/config"
	"github.com/gunjan01/data_pipeline/source/search"
	"github.com/sirupsen/logrus"
)

// CSVData is a struct to store CSV data.
type CSVData struct {
	Date  string  `json:"date"`
	Query string  `json:"query"`
	Count float64 `json:"count"`
}

// LoadCSV reads a CSV from the command line and loads it into the ES index.
func LoadCSV(path string) {
	csvfile, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	firstRow, err := bufio.NewReader(csvfile).ReadSlice('\n')
	if err != nil {
		panic(err)
	}
	// Skip the first row of the CSV. Read the remaining rows
	// and insert data into elastic index.
	_, err = csvfile.Seek(int64(len(firstRow)), io.SeekStart)
	if err != nil {
		panic(err)
	}

	client, err := search.New()
	if err != nil {
		panic(err)
	}

	var dataCSV CSVData
	ctx := context.Background()
	// Parse the file
	r := csv.NewReader(csvfile)

	for {
		// Read each record from csv
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		count, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			logrus.Errorf("Could not convert string into float")
		}

		dataCSV = CSVData{
			Date:  record[0],
			Query: record[1],
			Count: float64(count),
		}

		dataJSON, err := json.Marshal(dataCSV)
		js := string(dataJSON)

		_, er := client.Client.Index().
			Index(config.SearchIndex).
			Type("dimensions").
			BodyJson(js).
			Do(ctx)
		if er != nil {
			panic(er)
		}

		logrus.Infof("Insert operation successful")
	}
}
