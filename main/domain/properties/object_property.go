package properties

import "fmt"

type ObjectProperty struct {
	Name       string     `json:"name"`
	Type       Type       `json:"type"`
	Properties []Property `json:"properties"`
}

func (o *ObjectProperty) GetType() Type {
	return o.Type
}

func (o *ObjectProperty) GetName() string {
	return o.Name
}

func NewObjectProperty(name string, properties []Property) (*ObjectProperty, error) {
	err := validateObject(name, properties)
	if err != nil {
		return nil, err
	}
	return &ObjectProperty{Name: name, Type: Object, Properties: properties}, nil
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
