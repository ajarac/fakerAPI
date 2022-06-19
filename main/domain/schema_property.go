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

type SchemaProperty struct {
	Name       string            `json:"name"`
	Type       Type              `json:"type,omitempty"`
	Min        int               `json:"min,omitempty"`
	Max        int               `json:"max,omitempty"`
	Rate       float64           `json:"rate,omitempty"`
	Element    *SchemaProperty   `json:"element,omitempty"`
	RangeSize  []int             `json:"rangeSize,omitempty"`
	Properties []*SchemaProperty `json:"properties,omitempty" json:"properties"`
}
