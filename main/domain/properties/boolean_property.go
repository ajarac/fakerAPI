package properties

type BooleanProperty struct {
	AbstractProperty `bson:"metadata"`
	Rate             float32 `bson:"rate"`
}

func NewBooleanProperty(name string, rate float32) (*BooleanProperty, error) {
	err := validateBoolean(name, rate)
	if err != nil {
		return nil, err
	}
	return &BooleanProperty{
		AbstractProperty: newAbstractProperty(name, Boolean),
		Rate:             rate,
	}, nil
}

func validateBoolean(name string, rate float32) error {
	if 0 > rate || rate > 1 {
		return NewPropertyNotValid(name, "Rate should be between 0 and 1")
	}
	return nil
}
