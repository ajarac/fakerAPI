package in

import (
	"fakerAPI/main/application/out"
	"fakerAPI/main/domain"
)

type GetSchemasUseCase struct {
	storage out.SchemaStorage
}

func (u *GetSchemasUseCase) Get() []*domain.Schema {
	return u.storage.GetAll()
}

func NewGetSchemasUseCase(storage out.SchemaStorage) *GetSchemasUseCase {
	return &GetSchemasUseCase{storage: storage}
}
