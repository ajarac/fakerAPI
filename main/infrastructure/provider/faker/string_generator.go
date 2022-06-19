package faker

import (
	"fakerAPI/main/domain"
	"github.com/bxcodec/faker/v3"
)

func stringGenerator(property *domain.SchemaProperty) any {
	return faker.Word()
}
