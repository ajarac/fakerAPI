package faker

import (
	"fakerAPI/main/domain/properties"
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

func dateGenerator(property *properties.DateProperty, faker *gofakeit.Faker) time.Time {
	min := property.From
	max := property.To
	return faker.DateRange(min, max)
}
