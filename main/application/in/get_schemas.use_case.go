package in

import (
	"context"
	"fakerAPI/main/application/dto"
	"fakerAPI/main/application/out"
)

type GetSchemasUseCase struct {
	storage out.SchemaStorage
}

func (u *GetSchemasUseCase) Get(ctx context.Context) ([]*dto.BasicSchemaDTO, error) {
	return u.storage.GetAll(ctx)
}

func NewGetSchemasUseCase(storage out.SchemaStorage) *GetSchemasUseCase {
	return &GetSchemasUseCase{storage: storage}
}
