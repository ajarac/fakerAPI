package faker

import (
	"fakerAPI/main/domain"
	"fakerAPI/main/domain/properties"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"hash/fnv"
)

type ValueProvider struct {
}

func (provider *ValueProvider) Generate(schema *domain.Schema, valueId string) *domain.Value {
	value := make(domain.Value)
	key := fmt.Sprint(schema.Id, "-", valueId)
	seed := FNV64a(key)
	fmt.Println(seed)
	faker := gofakeit.New(seed)
	for _, prop := range schema.Properties {
		value[prop.GetName()] = provider.generateProperty(prop, faker)
	}
	return &value
}

func (provider *ValueProvider) generateProperty(property properties.Property, faker *gofakeit.Faker) any {
	switch property.GetType() {
	case properties.String:
		return stringGenerator(property.(*properties.StringProperty), faker)
	case properties.Number:
		return numberGenerator(property.(*properties.NumberProperty), faker)
	case properties.Boolean:
		return boolGenerator(property.(*properties.BooleanProperty), faker)
	case properties.Date:
		return dateGenerator(property.(*properties.DateProperty), faker)
	case properties.Object:
		return provider.objectGenerator(property.(*properties.ObjectProperty), faker)
	case properties.Array:
		return provider.arrayGenerator(property.(*properties.ArrayProperty), faker)
	case properties.Enum:
		return enumGenerator(property.(*properties.EnumProperty), faker)
	}
	return nil
}

func (provider *ValueProvider) objectGenerator(property *properties.ObjectProperty, faker *gofakeit.Faker) any {
	obj := make(map[string]interface{})
	for _, prop := range property.Properties {
		obj[prop.GetName()] = provider.generateProperty(prop, faker)
	}
	return obj
}

func (provider *ValueProvider) arrayGenerator(property *properties.ArrayProperty, faker *gofakeit.Faker) any {
	size := faker.Number(property.RangeSize[0], property.RangeSize[1]-1)
	list := make([]interface{}, size)
	for i := 0; i < size; i++ {
		list[i] = provider.generateProperty(property.Element, faker)
	}
	return list
}

func NewFakerValueProvider() *ValueProvider {
	return &ValueProvider{}
}

func FNV64a(text string) int64 {
	algorithm := fnv.New64()
	algorithm.Write([]byte(text))
	return int64(algorithm.Sum64())
}
