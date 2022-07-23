package mapper

import (
	"fakerAPI/main/domain/properties"
	"fmt"
	"time"
)

func BuildProperties(jsonProperties []JsonProperty) ([]properties.Property, error) {
	p := make([]properties.Property, len(jsonProperties))
	for i, property := range jsonProperties {
		prop, err := buildProperty(property)
		if err != nil {
			return nil, err
		}
		p[i] = prop
	}
	return p, nil
}

func buildProperty(p JsonProperty) (properties.Property, error) {
	switch p.Type {
	case properties.String:
		return properties.NewStringProperty(p.Name)
	case properties.Number:
		return properties.NewNumberProperty(p.Name, p.Min, p.Max)
	case properties.Boolean:
		return properties.NewBooleanProperty(p.Name, p.Rate)
	case properties.Date:
		if p.From.IsZero() {
			p.From = time.UnixMilli(0)
		}
		if p.To.IsZero() {
			p.To = time.Now()
		}
		return properties.NewDateProperty(p.Name, p.From, p.To)
	case properties.Object:
		buildProperties, err := BuildProperties(p.Properties)
		if err != nil {
			return nil, err
		}
		return properties.NewObjectProperty(p.Name, buildProperties)
	case properties.Array:
		element, err := buildProperty(*p.Element)
		if err != nil {
			return nil, err
		}
		return properties.NewArrayProperty(p.Name, p.RangeSize, element)
	case properties.Enum:
		return properties.NewEnumProperty(p.Name, p.Enums)
	default:
		return nil, properties.NewPropertyNotValid(p.Name, fmt.Sprintf("property %s is not valid", p.Type))
	}
}
