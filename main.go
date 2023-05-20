package main

import (
	"fmt"
	"os"
	"time"

	"github.com/griggsca91/fulltext_search_postgres/importer"
	"github.com/gocarina/gocsv"
)

const ()


var importers = []importer.Importer{
	// importer.NewPostgreSQLImporter(),
	importer.NewElasticSearchImporter(),
}

func main() {
	file, err := os.OpenFile("./cards.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, importer := range importers {
		if err := importer.Setup(); err != nil {
			panic(err)
		}
	}

	defer func() {
		for _, importer := range importers {
			if err := importer.Finish(); err != nil {
				panic(err)
			}
		}
	}()

	start := time.Now()


	var cards []importer.Card
	if err := gocsv.UnmarshalFile(file, &cards); err != nil { // Load clients from file
		panic(err)
	}

	for _, card := range cards {
		for _, importer := range importers {
			if err := importer.Upload(card); err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("total time", time.Since(start))
}
