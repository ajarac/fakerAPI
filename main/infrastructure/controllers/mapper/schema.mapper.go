package mapper

import (
	"fakerAPI/main/domain"
)

func BindToSchema(json *JsonSchema) *domain.Schema {
	schema := &domain.Schema{}

	return schema
}

func buildProperty(p *JsonProperty) *property.Property {
	switch p.Type {
	case string(domain.String):
		return &domain.String{
			Type:   property.STRING,
			Length: 0,
		}

	case property.Number:

	}
}
