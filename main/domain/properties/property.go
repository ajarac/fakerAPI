package properties

type Type string

const (
	String  Type = "string"
	Number       = "number"
	Boolean      = "boolean"
	Date         = "date"
	Object       = "object"
	Array        = "array"
	Enum         = "enum"
)

type Property interface {
	GetType() Type
	GetName() string
}
