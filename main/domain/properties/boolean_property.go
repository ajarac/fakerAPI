package properties

type BooleanProperty struct {
	Name string  `json:"name"`
	Type Type    `json:"type"`
	Rate float32 `json:"rate"`
}

func (b *BooleanProperty) GetType() Type {
	return b.Type
}

func (b *BooleanProperty) GetName() string {
	return b.Name
}

func NewBooleanProperty(name string, rate float32) (*BooleanProperty, error) {
	err := validateBoolean(name, rate)
	if err != nil {
		return nil, err
	}
	return &BooleanProperty{Name: name, Type: Boolean, Rate: rate}, nil
}

func validateBoolean(name string, rate float32) error {
	if 0 > rate || rate > 1 {
		return NewPropertyNotValid(name, "Rate should be between 0 and 1")
	}
	return nil
}
