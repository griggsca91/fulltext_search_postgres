package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/griggsca91/fulltext_search_postgres/importer"
	"github.com/gocarina/gocsv"
)

const ()


var importers = []importer.Importer{
	importer.NewPostgreSQLImporter(),
	importer.NewElasticSearchImporter(),
}

func main() {

	data, err := os.ReadFile("./cards.csv")
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("./cards.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, importer := range importers {
		importer.Setup();
	}



	reader := bytes.NewReader(data)

	csvReader := csv.NewReader(reader)
	isFirstRow := true
	headerMap := make(map[string]int)

	start := time.Now()


	var cards []importer.Card
	if err := gocsv.UnmarshalFile(file, &cards); err != nil { // Load clients from file
		panic(err)
	}


	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if isFirstRow {

			isFirstRow = false

			// Add mapping: Column/property name --> record index
			for i, v := range record {
				headerMap[v] = i
			}

			// Skip next code
			continue
		}
	}

	for _, importer := range importers {
		importer.Finish();
	}

	fmt.Println("total time", time.Since(start))
}
