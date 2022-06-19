package in

import (
	"fakerAPI/main/application/out"
	"fakerAPI/main/domain"
)

type GetValueUseCase struct {
	storage  out.SchemaStorage
	provider out.ValueProvider
}

func (u *GetValueUseCase) GetByName(name string) (*domain.Value, error) {
	schema, err := u.storage.GetByName(name)
	if err != nil {
		return nil, err
	}
	return u.provider.Generate(schema), nil
}

func NewGetValueUseCase(storage out.SchemaStorage, provider out.ValueProvider) *GetValueUseCase {
	return &GetValueUseCase{
		storage:  storage,
		provider: provider,
	}
}
