package faker

import (
	"fakerAPI/main/domain"
	"math/rand"
)

func numberGenerator(property *domain.SchemaProperty) any {
	return rand.Intn(100)
}
