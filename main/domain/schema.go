package domain

type Schema struct {
	Id         string              `json:"id"`
	Name       string              `json:"name"`
	Properties []*AbstractProperty `json:"properties"`
}

func NewSchema(id string, name string, properties []*AbstractProperty) *Schema {
	return &Schema{Id: id, Name: name, Properties: properties}
}
