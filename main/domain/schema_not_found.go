package domain

import "fmt"

type SchemaNotFound struct {
	msg string
}

func NewSchemaNotFoundByName(id string) *SchemaNotFound {
	return &SchemaNotFound{msg: fmt.Sprintf("Schema with id: %s not found", id)}
}

func (e *SchemaNotFound) Error() string {
	return e.msg
}
