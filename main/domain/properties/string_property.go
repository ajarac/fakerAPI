package properties

type StringProperty struct {
	AbstractProperty
}

func NewStringProperty(name string) (*StringProperty, error) {
	return &StringProperty{
		newAbstractProperty(name, String),
	}, nil
}
