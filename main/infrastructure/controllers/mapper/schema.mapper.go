package mapper

import (
	"fakerAPI/main/domain"
	"fakerAPI/main/domain/property"
)

func BindToSchema(json *JsonSchema) *domain.Schema {
	schema := &domain.Schema{}

	return schema
}

func buildProperty(p *JsonProperty) *property.Property {
	switch p.Type {
	case property.STRING:
		return &property.String{
			Type:   property.STRING,
			Length: 0,
		}
	}
}
