package faker

import (
	"fakerAPI/main/domain/properties"
	"math/rand"
)

func boolGenerator(property *properties.BooleanProperty) any {
	return rand.Float32() < property.Rate
}
