package in

import (
	"context"
	"fakerAPI/main/application/out"
	"fakerAPI/main/domain"
)

type GetListValueUseCase struct {
	storage  out.SchemaStorage
	provider out.ValueProvider
}

func (u *GetListValueUseCase) GetById(ctx context.Context, idOrName string, limit int) ([]*domain.Value, error) {
	schema, ok, err := u.storage.GetByIdOrByName(ctx, idOrName)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, domain.NewSchemaNotFoundByIdOrByName(idOrName)
	}
	if limit > 100 {
		limit = 100
	}
	var list = make([]*domain.Value, limit)
	for i := 0; i < limit; i++ {
		list[i] = u.provider.Generate(schema)
	}
	return list, nil
}

func NewGetListValueUseCase(storage out.SchemaStorage, provider out.ValueProvider) *GetListValueUseCase {
	return &GetListValueUseCase{storage: storage, provider: provider}
}
