package in

import "fakerAPI/main/application/out"

type DeleteSchemaUseCase struct {
	storage out.SchemaStorage
}

func (u *DeleteSchemaUseCase) Delete(id string) error {
	return u.storage.Delete(id)
}

func NewDeleteSchemaUseCase(storage out.SchemaStorage) *DeleteSchemaUseCase {
	return &DeleteSchemaUseCase{storage: storage}
}
