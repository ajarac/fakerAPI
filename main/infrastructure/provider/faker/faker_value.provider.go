package faker

import (
	"fakerAPI/main/domain"
	"math/rand"
)

type ValueProvider struct {
}

func (provider *ValueProvider) Generate(schema *domain.Schema) *domain.Value {
	value := make(domain.Value)
	for _, prop := range schema.Properties {
		value[prop.Name] = provider.generateProperty(prop)
	}
	return &value
}

func (provider *ValueProvider) generateProperty(property interface{}) any {

	switch v := property.(type) {
	case domain.PropertyString:
		return stringGenerator(&v)
	case domain.NumberProperty:
		return numberGenerator(&v)
	case domain.BooleanProperty:
		return boolGenerator(&v)
	case domain.DateProperty:
		return dateGenerator(&v)
	case domain.ObjectProperty:
		return provider.objectGenerator(&v)
	case domain.ArrayProperty:
		return provider.arrayGenerator(&v)
	}
	return nil
}

func (provider *ValueProvider) objectGenerator(property *domain.ObjectProperty) any {
	obj := make(map[string]interface{})
	for _, prop := range property.Properties {
		obj[prop.Name] = provider.generateProperty(prop)
	}
	return obj
}

func (provider *ValueProvider) arrayGenerator(property *domain.ArrayProperty) any {
	size := rand.Intn(property.RangeSize[1]-property.RangeSize[0]) + property.RangeSize[0]
	list := make([]interface{}, size)
	for i := 0; i < size; i++ {
		list[i] = provider.generateProperty(property.Element)
	}
	return list
}

func NewFakerValueProvider() *ValueProvider {
	return &ValueProvider{}
}
