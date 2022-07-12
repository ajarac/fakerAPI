package storage

import (
	"fakerAPI/main/domain"
	"fakerAPI/main/domain/properties"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func toDomain(mongoSchema *MongoSchema) *domain.Schema {
	return &domain.Schema{
		Id:         mongoSchema.Id.Hex(),
		Name:       mongoSchema.Name,
		Properties: toDomainProperties(mongoSchema.Properties),
	}
}

func toDomainProperties(mongoProperties []MongoSchemaProperty) []properties.Property {
	p := make([]properties.Property, len(mongoProperties))
	if mongoProperties != nil && len(mongoProperties) > 0 {
		for i, schemaProperty := range mongoProperties {
			p[i] = toDomainProperty(schemaProperty)
		}
	}
	return p
}

func toDomainProperty(mongoProperty MongoSchemaProperty) properties.Property {
	switch mongoProperty.Type {
	case properties.String:
		return properties.NewStringProperty(mongoProperty.Name)
	case properties.Number:
		return properties.NewNumberProperty(mongoProperty.Name, mongoProperty.Min, mongoProperty.Max)
	case properties.Boolean:
		return properties.NewBooleanProperty(mongoProperty.Name, mongoProperty.Rate)
	case properties.Date:
		return properties.NewDateProperty(mongoProperty.Name, mongoProperty.From, mongoProperty.To)
	case properties.Object:
		p := make([]properties.Property, len(mongoProperty.Properties))
		for i, objectProperty := range mongoProperty.Properties {
			domainProperty := toDomainProperty(objectProperty)
			p[i] = domainProperty
		}
		return properties.NewObjectProperty(mongoProperty.Name, p)
	case properties.Array:
		element := toDomainProperty(*mongoProperty.Element)
		return properties.NewArrayProperty(mongoProperty.Name, mongoProperty.RangeSize, element)
	}
	return nil
}

func fromDomain(schema *domain.Schema) (*MongoSchema, error) {
	id, err := primitive.ObjectIDFromHex(schema.Id)
	if err != nil {
		return nil, err
	}
	return &MongoSchema{
		Id:         id,
		UserId:     schema.UserId,
		Name:       schema.Name,
		Properties: fromDomainProperties(schema.Properties),
	}, nil
}

func fromDomainProperties(properties []properties.Property) []MongoSchemaProperty {
	mongoProperties := make([]MongoSchemaProperty, len(properties))
	if properties != nil && len(properties) > 0 {
		for i, schemaProperty := range properties {
			mongoProperties[i] = fromDomainProperty(schemaProperty)
		}
	}
	return mongoProperties
}

func fromDomainProperty(property properties.Property) MongoSchemaProperty {
	mongoProperty := MongoSchemaProperty{
		Type: property.GetType(),
		Name: property.GetName(),
	}
	switch property.GetType() {
	case properties.Number:
		mongoProperty.Min = property.(*properties.NumberProperty).Min
		mongoProperty.Max = property.(*properties.NumberProperty).Max
		break
	case properties.Boolean:
		mongoProperty.Rate = property.(*properties.BooleanProperty).Rate
		break
	case properties.Date:
		mongoProperty.From = property.(*properties.DateProperty).From
		mongoProperty.To = property.(*properties.DateProperty).To
		break
	case properties.Object:
		mongoProperties := make([]MongoSchemaProperty, len(property.(*properties.ObjectProperty).Properties))
		for i, objectProperty := range property.(*properties.ObjectProperty).Properties {
			mongoProperties[i] = fromDomainProperty(objectProperty)
		}
		mongoProperty.Properties = mongoProperties
		break
	case properties.Array:
		element := fromDomainProperty(property.(*properties.ArrayProperty).Element)
		mongoProperty.Element = &element
		mongoProperty.RangeSize = property.(*properties.ArrayProperty).RangeSize
		break
	}
	return mongoProperty
}
