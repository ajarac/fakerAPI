package in

import (
	"context"
	"fakerAPI/main/application/out"
	"fakerAPI/main/domain"
)

type GetSchemaUseCase struct {
	storage out.SchemaStorage
}

func (u *GetSchemaUseCase) GetById(ctx context.Context, id string) (*domain.Schema, bool, error) {
	return u.storage.GetById(ctx, id)
}

func NewGetSchemaUseCase(storage out.SchemaStorage) *GetSchemaUseCase {
	return &GetSchemaUseCase{storage: storage}
}
