package properties

type ArrayProperty struct {
	AbstractProperty `bson:"metadata"`
	RangeSize        [2]int   `bson:"rangeSize"`
	Element          Property `bson:"element"`
}

func NewArrayProperty(name string, rangeSize [2]int, element Property) *ArrayProperty {
	return &ArrayProperty{
		AbstractProperty: newAbstractProperty(name, Array),
		RangeSize:        rangeSize,
		Element:          element,
	}
}
