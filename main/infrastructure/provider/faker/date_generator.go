package faker

import (
	"fakerAPI/main/domain"
	"github.com/bxcodec/faker/v3"
)

func dateGenerator(property *domain.SchemaProperty) any {
	return faker.Date()
}
