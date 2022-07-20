package in

import (
	"context"
	"fakerAPI/main/application/out"
	"fakerAPI/main/domain"
)

type GetValueUseCase struct {
	storage  out.SchemaStorage
	provider out.ValueProvider
}

func (u *GetValueUseCase) GetById(ctx context.Context, id string) (*domain.Value, error) {
	schema, ok, err := u.storage.GetByIdOrByName(ctx, id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, domain.NewSchemaNotFoundByIdOrByName(id)
	}
	return u.provider.Generate(schema), nil
}

func NewGetValueUseCase(storage out.SchemaStorage, provider out.ValueProvider) *GetValueUseCase {
	return &GetValueUseCase{
		storage:  storage,
		provider: provider,
	}
}
