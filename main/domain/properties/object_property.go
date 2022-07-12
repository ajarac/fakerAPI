package properties

import "fmt"

type ObjectProperty struct {
	AbstractProperty `bson:"metadata"`
	Properties       []Property `bson:"properties"`
}

func NewObjectProperty(name string, properties []Property) (*ObjectProperty, error) {
	err := validateObject(name, properties)
	if err != nil {
		return nil, err
	}
	return &ObjectProperty{
		AbstractProperty: newAbstractProperty(name, Object),
		Properties:       properties,
	}, nil
}

func validateObject(objName string, props []Property) error {
	names := make(map[string]bool)
	for _, property := range props {
		name := property.GetName()
		if _, ok := names[name]; ok {
			return NewPropertyNotValid(objName, fmt.Sprintf("Property %s repeated", name))
		}
		names[name] = true
	}
	return nil
}
