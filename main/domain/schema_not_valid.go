package domain

type SchemaNotValid struct {
	msg string
}

func NewSchemaNotValid(msg string) *SchemaNotValid {
	return &SchemaNotValid{msg: msg}
}

func (e *SchemaNotValid) Error() string {
	return e.msg
}
