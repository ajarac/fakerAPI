package out

import "fakerAPI/main/domain"

type SchemaStorage interface {
	Create(schema *domain.Schema) (*domain.Schema, error)

	GetNextId() string

	GetById(id string) (*domain.Schema, bool, error)

	GetAll() ([]*domain.Schema, error)

	Delete(id string) error
}
