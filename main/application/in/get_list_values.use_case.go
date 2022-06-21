package in

import (
	"fakerAPI/main/application/out"
	"fakerAPI/main/domain"
)

type GetListValueUseCase struct {
	storage  out.SchemaStorage
	provider out.ValueProvider
}

func (u *GetListValueUseCase) GetById(id string, limit int) ([]*domain.Value, error) {
	schema, ok, err := u.storage.GetById(id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, domain.NewSchemaNotFoundByName(id)
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
