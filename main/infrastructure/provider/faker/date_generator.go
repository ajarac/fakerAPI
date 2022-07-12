package faker

import (
	"fakerAPI/main/domain/properties"
	"math/rand"
	"time"
)

func dateGenerator(property *properties.DateProperty) time.Time {
	min := property.From.Unix()
	max := property.To.Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
