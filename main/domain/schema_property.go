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
	Name       string            `json:"name" bson:"name"`
	Type       Type              `json:"type,omitempty" bson:"type"`
	Min        int               `json:"min,omitempty" bson:"min,omitempty"`
	Max        int               `json:"max,omitempty" bson:"max,omitempty"`
	Rate       float64           `json:"rate,omitempty" bson:"rate,omitempty"`
	Element    *SchemaProperty   `json:"element,omitempty" bson:"element,omitempty"`
	RangeSize  []int             `json:"rangeSize,omitempty" bson:"rangeSize,omitempty"`
	Properties []*SchemaProperty `json:"properties,omitempty" bson:"properties,omitempty"`
}
