package properties

type StringProperty struct {
	AbstractProperty `bson:"metadata"`
}

func NewStringProperty(name string) (*StringProperty, error) {
	return &StringProperty{
		newAbstractProperty(name, String),
	}, nil
}
