package in

import (
	"fakerAPI/main/application/out"
	"fakerAPI/main/domain"
)

type CreateSchemaUseCase struct {
	storage out.SchemaStorage
}

func (s *CreateSchemaUseCase) Create(schema *domain.Schema) error {
	return s.storage.Create(schema)
}

func NewCreateSchemaUseCase(storage out.SchemaStorage) *CreateSchemaUseCase {
	return &CreateSchemaUseCase{storage: storage}
}
