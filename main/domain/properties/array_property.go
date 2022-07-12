package properties

type ArrayProperty struct {
	AbstractProperty `bson:"metadata"`
	RangeSize        [2]int   `bson:"rangeSize"`
	Element          Property `bson:"element"`
}

func NewArrayProperty(name string, rangeSize [2]int, element Property) (*ArrayProperty, error) {
	err := validateArray(name, rangeSize)
	if err != nil {
		return nil, err
	}
	return &ArrayProperty{
		AbstractProperty: newAbstractProperty(name, Array),
		RangeSize:        rangeSize,
		Element:          element,
	}, nil
}

func validateArray(name string, rangeSize [2]int) error {
	if rangeSize[0] > rangeSize[1] {
		return NewPropertyNotValid(name, "Range Size invalid, the first number should be less than the second number")
	}
	return nil
}
