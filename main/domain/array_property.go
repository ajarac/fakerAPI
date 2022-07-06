package domain

type ArrayProperty struct {
	AbstractProperty
	RangeSize [2]int
	Element   *AbstractProperty
}

func (a *ArrayProperty) GetType() Type {
	return Array
}

func NewArrayProperty(rangeSize [2]int, element *AbstractProperty) *ArrayProperty {
	return &ArrayProperty{
		RangeSize: rangeSize,
		Element:   element,
	}
}
