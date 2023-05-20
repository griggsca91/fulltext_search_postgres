package clients

import (
	"bytes"
	"strconv"
	"sync/atomic"

	"github.com/dustin/go-humanize"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/griggsca91/fulltext_search_postgres/importer"

	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Get the certificate with this command
//docker cp fulltext_search_postgres-elasticsearch01-1:/usr/share/elasticsearch/config/certs/http_ca.crt .

type ElasticClient struct {
	bi esutil.BulkIndexer
	countSuccessful uint64

}

func connect() (*elasticsearch.Client, error) {
	cert, err := os.ReadFile("./http_ca.crt")
	if err != nil {
		return nil, err
	}

	cfg := elasticsearch.Config{
		CACert: cert,
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "example1",
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	res, err := es.Info()
	if err != nil {
		return nil, err
	}

	fmt.Println("res", res)


	return es, nil
}

func createBulkIndexer(es *elasticsearch.Client, indexName string) (esutil.BulkIndexer, error) {
		// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//
	// Create the BulkIndexer
	//
	// NOTE: For optimal performance, consider using a third-party JSON decoding package.
	//       See an example in the "benchmarks" folder.
	//
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         indexName,        // The default index name
		Client:        es,               // The Elasticsearch client
		NumWorkers:    10,       // The number of worker goroutines
		FlushBytes:    int(5e+6),  // The flush threshold in bytes
		FlushInterval: 30 * time.Second, // The periodic flush interval
	})
	return bi, err

}

func (e ElasticClient) Test() {
	_, err := connect()
	if err != nil {
		panic(err)
	}
}

func (e ElasticClient) Setup() error {
	client, err := connect()
	bi, err := createBulkIndexer(client, "cards")
	e.bi = bi

	return err
}

func (e ElasticClient) Upload(card importer.Card) error {
	return e.bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				// Action field configures the operation to perform (index, create, delete, update)
				Action: "index",

				// DocumentID is the (optional) document ID
				DocumentID: strconv.Itoa(card.ID),

				// Body is an `io.Reader` with the payload
				Body: bytes.NewReader(data),

				// OnSuccess is called for each successful operation
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},

				// OnFailure is called for each failed operation
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
}


func (e ElasticClient) Finish() error {
		if err := e.bi.Close(context.Background()); err != nil {
		return err
	}
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	biStats := e.bi.Stats()

	// Report the results: number of indexed docs, number of errors, duration, indexing rate
	//
	log.Println(strings.Repeat("▔", 65))
		if biStats.NumFailed > 0 {
		log.Fatalf(
			"Indexed [%s] documents with [%s] errors in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			humanize.Comma(int64(biStats.NumFailed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	} else {
		log.Printf(
			"Sucessfuly indexed [%s] documents in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	}

	return nil

}
