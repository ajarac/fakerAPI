package faker

import (
	"fakerAPI/main/domain"
	"math/rand"
)

type mapGenerator = map[domain.Type]func(property *domain.SchemaProperty) any

type ValueProvider struct {
	delegate mapGenerator
}

func (f *ValueProvider) Generate(schema *domain.Schema) *domain.Value {
	value := make(domain.Value)
	for _, prop := range schema.Properties {
		value[prop.Name] = f.generateProperty(prop)
	}
	return &value
}

func (f *ValueProvider) generateProperty(property *domain.SchemaProperty) any {
	return f.delegate[property.Type](property)
}

func NewFakerValueProvider() *ValueProvider {
	valueProvider := ValueProvider{}
	valueProvider.delegate = mapGenerator{
		domain.String:  stringGenerator,
		domain.Number:  numberGenerator,
		domain.Boolean: boolGenerator,
		domain.Date:    dateGenerator,
		domain.Object: func(property *domain.SchemaProperty) any {
			obj := make(map[string]interface{})
			for _, prop := range property.Properties {
				obj[prop.Name] = valueProvider.generateProperty(prop)
			}
			return obj
		},
		domain.Array: func(property *domain.SchemaProperty) any {
			size := rand.Intn(property.RangeSize[1]-property.RangeSize[0]) + property.RangeSize[0]
			list := make([]interface{}, size)
			for i := 0; i < size; i++ {
				list[i] = valueProvider.generateProperty(property.Element)
			}
			return list
		},
	}
	return &valueProvider
}
