package domain

type Schema struct {
	Name       string            `json:"name"`
	Properties []*SchemaProperty `json:"properties"`
}
