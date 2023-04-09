package out

import "fakerAPI/main/domain"

type ValueProvider interface {
	Generate(schema *domain.Schema, valueId string) *domain.Value
}
