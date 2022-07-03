package in

import (
	"context"
	"fakerAPI/main/application/out"
	"fakerAPI/main/domain"
)

type GetSchemasUseCase struct {
	storage out.SchemaStorage
}

func (u *GetSchemasUseCase) Get(ctx context.Context) ([]*domain.Schema, error) {
	return u.storage.GetAll(ctx)
}

func NewGetSchemasUseCase(storage out.SchemaStorage) *GetSchemasUseCase {
	return &GetSchemasUseCase{storage: storage}
}
