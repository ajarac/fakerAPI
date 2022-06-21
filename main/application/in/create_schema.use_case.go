package in

import (
	"fakerAPI/main/application/out"
	"fakerAPI/main/domain"
)

type CreateSchemaUseCase struct {
	storage out.SchemaStorage
}

func (s *CreateSchemaUseCase) Create(name string, properties []*domain.SchemaProperty) (*domain.Schema, error) {
	schema := domain.NewSchema(s.storage.GetNextId(), name, properties)
	return s.storage.Create(schema)
}

func NewCreateSchemaUseCase(storage out.SchemaStorage) *CreateSchemaUseCase {
	return &CreateSchemaUseCase{storage: storage}
}
