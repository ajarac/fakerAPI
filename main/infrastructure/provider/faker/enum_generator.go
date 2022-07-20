package faker

import (
	"fakerAPI/main/domain/properties"
	"math/rand"
)

func enumGenerator(property *properties.EnumProperty) string {
	randomIndex := rand.Intn(len(property.Enums))
	return property.Enums[randomIndex]
}
