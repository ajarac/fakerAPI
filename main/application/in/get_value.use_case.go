package in

import (
	"fakerAPI/main/application/out"
	"fakerAPI/main/domain"
)

type GetValueUseCase struct {
	storage  out.SchemaStorage
	provider out.ValueProvider
}

func (u *GetValueUseCase) GetById(id string) (*domain.Value, error) {
	schema, ok, err := u.storage.GetById(id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, domain.NewSchemaNotFoundByName(id)
	}
	return u.provider.Generate(schema), nil
}

func NewGetValueUseCase(storage out.SchemaStorage, provider out.ValueProvider) *GetValueUseCase {
	return &GetValueUseCase{
		storage:  storage,
		provider: provider,
	}
}
