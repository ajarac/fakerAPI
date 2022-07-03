package in

import (
	"context"
	"fakerAPI/main/application/out"
	"fakerAPI/main/domain"
)

type CreateSchemaUseCase struct {
	storage out.SchemaStorage
}

func (s *CreateSchemaUseCase) Create(context context.Context, name string, properties []*domain.SchemaProperty) (*domain.Schema, error) {
	schema := domain.NewSchema(s.storage.GetNextId(), name, properties)
	return s.storage.Create(context, schema)
}

func NewCreateSchemaUseCase(storage out.SchemaStorage) *CreateSchemaUseCase {
	return &CreateSchemaUseCase{storage: storage}
}
