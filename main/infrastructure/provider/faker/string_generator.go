package faker

import (
	"fakerAPI/main/domain/properties"
	"github.com/bxcodec/faker/v3"
)

func stringGenerator(property *properties.StringProperty) any {
	return faker.Word()
}
