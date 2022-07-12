package properties

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
	GetName() string
}

type AbstractProperty struct {
	Name string
	Type Type
}

func (p *AbstractProperty) GetType() Type {
	return p.Type
}
func (p *AbstractProperty) GetName() string {
	return p.Name
}

func newAbstractProperty(name string, t Type) AbstractProperty {
	return AbstractProperty{
		Name: name,
		Type: t,
	}
}
