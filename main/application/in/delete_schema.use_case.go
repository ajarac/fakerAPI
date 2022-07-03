package in

import (
	"context"
	"fakerAPI/main/application/out"
)

type DeleteSchemaUseCase struct {
	storage out.SchemaStorage
}

func (u *DeleteSchemaUseCase) Delete(ctx context.Context, id string) error {
	return u.storage.Delete(ctx, id)
}

func NewDeleteSchemaUseCase(storage out.SchemaStorage) *DeleteSchemaUseCase {
	return &DeleteSchemaUseCase{storage: storage}
}
