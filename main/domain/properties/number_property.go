package properties

type NumberProperty struct {
	Name string `json:"name"`
	Type Type   `json:"type"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
}

func (n *NumberProperty) GetType() Type {
	return n.Type
}

func (n *NumberProperty) GetName() string {
	return n.Name
}

func NewNumberProperty(name string, min int, max int) (*NumberProperty, error) {
	err := validateNumber(name, min, max)
	if err != nil {
		return nil, err
	}
	return &NumberProperty{Name: name, Type: Number, Min: min, Max: max}, nil
}

func validateNumber(name string, min int, max int) error {
	if min > max {
		return NewPropertyNotValid(name, "Min value should be less than Max value")
	}
	return nil
}
