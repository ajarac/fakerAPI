package domain

import "fmt"

type SchemaAlreadyExists struct {
	msg string
}

func NewSchemaAlreadyExistsByName(name string) *SchemaAlreadyExists {
	return &SchemaAlreadyExists{msg: fmt.Sprintf("Schema %s already exists", name)}
}

func (e *SchemaAlreadyExists) Error() string {
	return e.msg
}
