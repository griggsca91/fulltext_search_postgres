package importer

import "github.com/jackc/pgx/v5"
import "context"

type PostgreSQLImporter struct {
	conn *pgx.Conn
}

// Finish implements Importer
func (p *PostgreSQLImporter) Finish() error {
	return p.conn.Close(context.Background())
}

// Setup implements Importer
func (p *PostgreSQLImporter) Setup() error {
	conn, err := pgx.Connect(context.Background(), "host=localhost port=5432 user=postgres password=example sslmode=disable")
	if err != nil {
		return err
	}

	p.conn = conn
	return nil
}

// Upload implements Importer
func (i *PostgreSQLImporter) Upload(card Card) error {
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

	_, err := i.conn.Exec(context.Background(), insertQuery,
			card.ID,
			card.Artist,
			card.Name,
			card.Colors,
			card.Defense,
			card.FlavorText,
			card.Life,
			card.ManaCost,
			card.Keywords,
			card.Number,
			card.Text,
			card.Power,
			card.Toughness,
			card.Type,
			card.Types,
			card.Subtypes,
			card.Supertypes,
		)

	return err
}

func NewPostgreSQLImporter() *PostgreSQLImporter {
	return &PostgreSQLImporter{}
}

var _ Importer = &PostgreSQLImporter{}
