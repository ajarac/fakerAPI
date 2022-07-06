package domain

type ObjectProperty struct {
	AbstractProperty
	Properties []*AbstractProperty
}

func (o *ObjectProperty) GetType() Type {
	return Object
}

func NewObjectProperty(properties []*AbstractProperty) *ObjectProperty {
	return &ObjectProperty{
		Properties: properties,
	}
}
