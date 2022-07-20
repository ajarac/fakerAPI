package domain

import "fmt"

type SchemaNotFound struct {
	msg string
}

func NewSchemaNotFoundByIdOrByName(id string) *SchemaNotFound {
	return &SchemaNotFound{msg: fmt.Sprintf("Schema with id or name: %s not found", id)}
}

func (e *SchemaNotFound) Error() string {
	return e.msg
}
