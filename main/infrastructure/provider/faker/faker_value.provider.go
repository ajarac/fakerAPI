package faker

import (
	"fakerAPI/main/domain"
	"fakerAPI/main/domain/properties"
	"math/rand"
)

type ValueProvider struct {
}

func (provider *ValueProvider) Generate(schema *domain.Schema) *domain.Value {
	value := make(domain.Value)
	for _, prop := range schema.Properties {
		value[prop.GetName()] = provider.generateProperty(prop)
	}
	return &value
}

func (provider *ValueProvider) generateProperty(property properties.Property) any {
	switch property.GetType() {
	case properties.String:
		return stringGenerator(property.(*properties.StringProperty))
	case properties.Number:
		return numberGenerator(property.(*properties.NumberProperty))
	case properties.Boolean:
		return boolGenerator(property.(*properties.BooleanProperty))
	case properties.Date:
		return dateGenerator(property.(*properties.DateProperty))
	case properties.Object:
		return provider.objectGenerator(property.(*properties.ObjectProperty))
	case properties.Array:
		return provider.arrayGenerator(property.(*properties.ArrayProperty))
	case properties.Enum:
		return enumGenerator(property.(*properties.EnumProperty))
	}
	return nil
}

func (provider *ValueProvider) objectGenerator(property *properties.ObjectProperty) any {
	obj := make(map[string]interface{})
	for _, prop := range property.Properties {
		obj[prop.GetName()] = provider.generateProperty(prop)
	}
	return obj
}

func (provider *ValueProvider) arrayGenerator(property *properties.ArrayProperty) any {
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
