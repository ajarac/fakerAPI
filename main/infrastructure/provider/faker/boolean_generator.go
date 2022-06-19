package faker

import (
	"fakerAPI/main/domain"
	"math/rand"
)

func boolGenerator(property *domain.SchemaProperty) any {
	return rand.Intn(2) == 1
}
