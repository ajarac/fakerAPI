package out

import "fakerAPI/main/domain"

type SchemaStorage interface {
	Create(schema *domain.Schema) error

	GetByName(name string) (*domain.Schema, error)

	GetAll() []*domain.Schema

	Delete(name string) error
}
