package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

const ()

func main() {
	data, err := os.ReadFile("../cards.csv")
	if err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(context.Background(), "host=localhost port=5432 user=postgres password=example sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	reader := bytes.NewReader(data)

	csvReader := csv.NewReader(reader)
	isFirstRow := true
	headerMap := make(map[string]int)

	start := time.Now()

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

		insertQuery := `insert into
cards(
	id,
	artist,
	name,
	colors,
	defense,
	flavorText,
	life,
	manaCost,
	keywords,
	number,
	text,
	power,
	toughness,
	type,
	types,
	subtypes,
	supertypes
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12,$13, $14, $15, $16, $17)`

		_, err = conn.Exec(context.Background(), insertQuery,
			record[headerMap["id"]],
			record[headerMap["artist"]],
			record[headerMap["name"]],
			record[headerMap["colors"]],
			record[headerMap["defense"]],
			record[headerMap["flavorText"]],
			record[headerMap["life"]],
			record[headerMap["manaCost"]],
			record[headerMap["keywords"]],
			record[headerMap["number"]],
			record[headerMap["text"]],
			record[headerMap["power"]],
			record[headerMap["toughness"]],
			record[headerMap["type"]],
			record[headerMap["types"]],
			record[headerMap["subtypes"]],
			record[headerMap["supertypes"]],
		)

		if err != nil {
			panic(err)
		}

	}

	fmt.Println("total time", time.Since(start))
}
