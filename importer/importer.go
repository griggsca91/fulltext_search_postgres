package importer

type Importer interface {
	Setup() error
	Upload(Card) error
	Finish() error
}

type Card struct {
	ID string  `json:"ID" csv:"id"`
	Artist string  `json:"artist" csv:"artist"`
	Name string  `json:"name" csv:"name"`
	Colors string  `json:"colors" csv:"colors"`
	Defense string  `json:"defense" csv:"defense"`
	FlavorText string  `json:"flavorText" csv:"flavorText"`
	Life string  `json:"life" csv:"life"`
	ManaCost string  `json:"manaCost" csv:"manaCost"`
	Keywords string  `json:"keywords" csv:"keywords"`
	Number string  `json:"number" csv:"number"`
	Text string  `json:"text" csv:"text"`
	Power string  `json:"power" csv:"power"`
	Toughness string  `json:"toughness" csv:"toughness"`
	Type string  `json:"type" csv:"type"`
	Types string  `json:"types" csv:"types"`
	Subtypes string  `json:"subtypes" csv:"subtypes"`
	Supertypes string  `json:"supertypes" csv:"supertypes"`
}
