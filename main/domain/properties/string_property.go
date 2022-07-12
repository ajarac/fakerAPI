package properties

type StringProperty struct {
	AbstractProperty `bson:"metadata"`
}

func NewStringProperty(name string) *StringProperty {
	return &StringProperty{
		newAbstractProperty(name, String),
	}
}
