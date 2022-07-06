package domain

type Type string

const (
	String  Type = "string"
	Number       = "number"
	Boolean      = "boolean"
	Date         = "date"
	Object       = "object"
	Array        = "array"
)

type Property interface {
	GetType() Type
}

type AbstractProperty struct {
	Name string
}

func (p *AbstractProperty) GetType() Type {
	return ""
}
