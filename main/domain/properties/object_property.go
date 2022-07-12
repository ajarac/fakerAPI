package properties

type ObjectProperty struct {
	AbstractProperty `bson:"metadata"`
	Properties       []Property `bson:"properties"`
}

func NewObjectProperty(name string, properties []Property) *ObjectProperty {
	return &ObjectProperty{
		AbstractProperty: newAbstractProperty(name, Object),
		Properties:       properties,
	}
}
