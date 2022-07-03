package out

import (
	"context"
	"fakerAPI/main/domain"
)

type SchemaStorage interface {
	Create(context context.Context, schema *domain.Schema) (*domain.Schema, error)

	GetNextId() string

	GetById(context context.Context, id string) (*domain.Schema, bool, error)

	GetAll(context context.Context) ([]*domain.Schema, error)

	Delete(context context.Context, id string) error
}
