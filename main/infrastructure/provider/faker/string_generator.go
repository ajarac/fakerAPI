package faker

import (
	"fakerAPI/main/domain/properties"
	"github.com/brianvoe/gofakeit/v6"
)

func stringGenerator(_ *properties.StringProperty, faker *gofakeit.Faker) any {
	return faker.Word()
}
