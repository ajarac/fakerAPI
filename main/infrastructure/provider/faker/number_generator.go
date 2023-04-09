package faker

import (
	"fakerAPI/main/domain/properties"
	"github.com/brianvoe/gofakeit/v6"
)

func numberGenerator(property *properties.NumberProperty, faker *gofakeit.Faker) any {
	max := property.Max
	min := property.Min
	return faker.Number(min, max)
}
