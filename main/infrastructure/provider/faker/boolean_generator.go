package faker

import (
	"fakerAPI/main/domain"
	"math/rand"
)

func boolGenerator(property *domain.BooleanProperty) any {
	return rand.Float32() < property.Rate
}
