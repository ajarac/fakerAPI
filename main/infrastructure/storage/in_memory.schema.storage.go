package storage

import (
	"fakerAPI/main/domain"
)

type InMemorySchemaStorage struct {
	schemas map[string]*domain.Schema
}

func NewInMemorySchemaStorage() *InMemorySchemaStorage {
	return &InMemorySchemaStorage{schemas: make(map[string]*domain.Schema)}
}

func (i *InMemorySchemaStorage) Create(schema *domain.Schema) error {
	i.schemas[schema.Name] = schema
	return nil
}

func (i *InMemorySchemaStorage) GetByName(name string) (*domain.Schema, error) {
	schema, ok := i.schemas[name]
	if ok == false {
		return nil, domain.NewSchemaNotFoundByName(name)
	}
	return schema, nil
}

func (i *InMemorySchemaStorage) GetAll() []*domain.Schema {
	var list []*domain.Schema
	for _, schema := range i.schemas {
		list = append(list, schema)
	}
	return list
}

func (i *InMemorySchemaStorage) Delete(name string) error {
	delete(i.schemas, name)
	return nil
}
