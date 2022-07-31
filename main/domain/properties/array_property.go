package properties

type ArrayProperty struct {
	Name      string   `json:"name"`
	Type      Type     `json:"type"`
	RangeSize [2]int   `json:"rangeSize"`
	Element   Property `json:"element"`
}

func (a *ArrayProperty) GetType() Type {
	return a.Type
}

func (a *ArrayProperty) GetName() string {
	return a.Name
}

func NewArrayProperty(name string, rangeSize [2]int, element Property) (*ArrayProperty, error) {
	err := validateArray(name, rangeSize)
	if err != nil {
		return nil, err
	}
	return &ArrayProperty{Name: name, Type: Array, RangeSize: rangeSize, Element: element}, nil
}

func validateArray(name string, rangeSize [2]int) error {
	if rangeSize[0] > rangeSize[1] {
		return NewPropertyNotValid(name, "Range Size invalid, the first number should be less than the second number")
	}
	return nil
}
