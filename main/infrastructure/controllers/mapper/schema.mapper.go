package mapper

import "fakerAPI/main/domain/properties"

func BuildProperties(jsonProperties []JsonProperty) []properties.Property {
	p := make([]properties.Property, len(jsonProperties))
	for i, property := range jsonProperties {
		prop := buildProperty(property)
		p[i] = prop
	}
	return p
}

func buildProperty(p JsonProperty) properties.Property {
	switch p.Type {
	case properties.String:
		return properties.NewStringProperty(p.Name)
	case properties.Number:
		return properties.NewNumberProperty(p.Name, p.Min, p.Max)
	case properties.Boolean:
		return properties.NewBooleanProperty(p.Name, p.Rate)
	case properties.Date:
		return properties.NewDateProperty(p.Name, p.From, p.To)
	case properties.Object:
		buildProperties := BuildProperties(p.Properties)
		return properties.NewObjectProperty(p.Name, buildProperties)
	case properties.Array:
		return properties.NewArrayProperty(p.Name, p.RangeSize, buildProperty(*p.Element))
	}
	return nil
}
