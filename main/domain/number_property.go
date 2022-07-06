package domain

type NumberProperty struct {
	AbstractProperty
	Min int
	Max int
}

func (n *NumberProperty) GetType() Type {
	return Number
}

func NewNumberProperty(min int, max int) *NumberProperty {
	return &NumberProperty{
		Min: min,
		Max: max,
	}
}
