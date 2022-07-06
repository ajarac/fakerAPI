package faker

import (
	"fakerAPI/main/domain"
	"github.com/bxcodec/faker/v3"
)

func stringGenerator(property *domain.PropertyString) any {
	return faker.Word()
}
