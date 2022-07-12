package domain

import "fakerAPI/main/domain/properties"

type Schema struct {
	Id         string
	Name       string
	Properties []properties.Property
}

func NewSchema(id string, name string, properties []properties.Property) *Schema {
	return &Schema{Id: id, Name: name, Properties: properties}
}
