package in

import (
	"context"
	"fakerAPI/main/application/out"
	"fakerAPI/main/domain"
	"fakerAPI/main/domain/properties"
)

type CreateSchemaUseCase struct {
	storage out.SchemaStorage
}

func (s *CreateSchemaUseCase) Create(context context.Context, name string, properties []properties.Property) (*domain.Schema, error) {
	_, exists, err := s.storage.GetByName(context, name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, domain.NewSchemaAlreadyExistsByName(name)
	}

	schema, err := domain.NewSchema(s.storage.GetNextId(), name, properties)
	if err != nil {
		return nil, err
	}
	return s.storage.Create(context, schema)
}

func NewCreateSchemaUseCase(storage out.SchemaStorage) *CreateSchemaUseCase {
	return &CreateSchemaUseCase{storage: storage}
}
