package in

import (
	"fakerAPI/main/application/out"
	"fakerAPI/main/domain"
)

type GetListValueUseCase struct {
	storage  out.SchemaStorage
	provider out.ValueProvider
}

func (u *GetListValueUseCase) GetByNameAndLimit(name string, limit int) ([]*domain.Value, error) {
	schema, err := u.storage.GetByName(name)
	if err != nil {
		return nil, err
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
