package faker

import (
	"fakerAPI/main/domain/properties"
	"github.com/brianvoe/gofakeit/v6"
)

func enumGenerator(property *properties.EnumProperty, faker *gofakeit.Faker) string {
	randomIndex := faker.Number(0, len(property.Enums)-1)
	return property.Enums[randomIndex]
}
