package properties

type NumberProperty struct {
	AbstractProperty
	Min int
	Max int
}

func NewNumberProperty(name string, min int, max int) (*NumberProperty, error) {
	err := validateNumber(name, min, max)
	if err != nil {
		return nil, err
	}
	return &NumberProperty{
		AbstractProperty: newAbstractProperty(name, Number),
		Min:              min,
		Max:              max,
	}, nil
}

func validateNumber(name string, min int, max int) error {
	if min > max {
		return NewPropertyNotValid(name, "Min value should be less than Max value")
	}
	return nil
}
