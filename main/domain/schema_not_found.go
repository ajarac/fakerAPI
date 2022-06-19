package domain

import "fmt"

type SchemaNotFound struct {
	msg string
}

func NewSchemaNotFoundByName(name string) *SchemaNotFound {
	return &SchemaNotFound{msg: fmt.Sprintf("Schema with name: %s not found", name)}
}

func (e *SchemaNotFound) Error() string {
	return e.msg
}
