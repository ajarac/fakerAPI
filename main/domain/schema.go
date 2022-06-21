package domain

type Schema struct {
	Id         string            `json:"id" bson:"_id"`
	Name       string            `json:"name" bson:"name"`
	Properties []*SchemaProperty `json:"properties" bson:"properties"`
}

func NewSchema(id string, name string, properties []*SchemaProperty) *Schema {
	return &Schema{Id: id, Name: name, Properties: properties}
}
