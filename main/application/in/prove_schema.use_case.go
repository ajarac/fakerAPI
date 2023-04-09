package in

import (
	"fakerAPI/main/application/out"
	"fakerAPI/main/domain"
	"fakerAPI/main/domain/properties"
)

type ProveSchemaUseCase struct {
	provider out.ValueProvider
}

func (u *ProveSchemaUseCase) Prove(properties []properties.Property) (*domain.Value, error) {
	schema, err := domain.NewSchema("", "testing", properties)
	if err != nil {
		return nil, err
	}
	return u.provider.Generate(schema, "0"), nil
}

func NewProveSchemaUseCase(provider out.ValueProvider) *ProveSchemaUseCase {
	return &ProveSchemaUseCase{provider: provider}
}
