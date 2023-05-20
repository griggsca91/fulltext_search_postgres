package importer

type ElasticSearchImporter struct {
}

// Finish implements Importer
func (*ElasticSearchImporter) Finish() error {
	panic("unimplemented")
}

// Setup implements Importer
func (*ElasticSearchImporter) Setup() error {
	panic("unimplemented")
}

// Upload implements Importer
func (*ElasticSearchImporter) Upload(Card) error {
	panic("unimplemented")
}

func NewElasticSearchImporter() *ElasticSearchImporter {
	return &ElasticSearchImporter{}
}

var _ Importer = &ElasticSearchImporter{}
