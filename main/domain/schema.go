package domain

import (
	"fakerAPI/main/domain/properties"
	"fmt"
)

type Schema struct {
	Id         string
	Name       string
	Properties []properties.Property
}

func NewSchema(id string, name string, properties []properties.Property) (*Schema, error) {
	err := validateSchema(properties)
	if err != nil {
		return nil, err
	}
	return &Schema{Id: id, Name: name, Properties: properties}, nil
}

func validateSchema(props []properties.Property) error {
	names := make(map[string]bool)
	for _, property := range props {
		name := property.GetName()
		if _, ok := names[name]; ok {
			return NewSchemaNotValid(fmt.Sprintf("Property %s repeated", name))
		}
		names[name] = true
	}
	return nil
}
