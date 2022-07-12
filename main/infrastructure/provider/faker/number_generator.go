package faker

import (
	"fakerAPI/main/domain/properties"
	"math/rand"
)

func numberGenerator(property *properties.NumberProperty) any {
	return rand.Intn(100)
}
