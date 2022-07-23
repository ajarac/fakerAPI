package faker

import (
	"fakerAPI/main/domain/properties"
	"math/rand"
)

func numberGenerator(property *properties.NumberProperty) any {
	max := property.Max
	min := property.Min
	return rand.Intn(max-min+1) + min
}
