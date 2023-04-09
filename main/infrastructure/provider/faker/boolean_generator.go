package faker

import (
	"fakerAPI/main/domain/properties"
	"github.com/brianvoe/gofakeit/v6"
)

func boolGenerator(property *properties.BooleanProperty, faker *gofakeit.Faker) any {
	return faker.Number(0, 100) < property.Rate
}
