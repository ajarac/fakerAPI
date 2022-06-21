package storage

import (
	"fakerAPI/main/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func toDomain(mongoSchema *MongoSchema) *domain.Schema {
	return &domain.Schema{
		Id:         mongoSchema.Id.Hex(),
		Name:       mongoSchema.Name,
		Properties: toDomainProperties(mongoSchema.Properties),
	}
}

func toDomainProperties(mongoProperties []*MongoSchemaProperty) []*domain.SchemaProperty {
	var properties []*domain.SchemaProperty
	if mongoProperties != nil && len(mongoProperties) > 0 {
		for _, schemaProperty := range mongoProperties {
			properties = append(properties, toDomainProperty(schemaProperty))
		}
	}
	return properties
}

func toDomainProperty(mongoProperty *MongoSchemaProperty) *domain.SchemaProperty {
	if mongoProperty == nil {
		return nil
	}
	return &domain.SchemaProperty{
		Name:       mongoProperty.Name,
		Type:       mongoProperty.Type,
		Min:        mongoProperty.Min,
		Max:        mongoProperty.Max,
		Rate:       mongoProperty.Rate,
		Element:    toDomainProperty(mongoProperty.Element),
		RangeSize:  mongoProperty.RangeSize,
		Properties: toDomainProperties(mongoProperty.Properties),
	}
}

func fromDomain(schema *domain.Schema) (*MongoSchema, error) {
	id, err := primitive.ObjectIDFromHex(schema.Id)
	if err != nil {
		return nil, err
	}
	return &MongoSchema{
		Id:         id,
		Name:       schema.Name,
		Properties: fromDomainProperties(schema.Properties),
	}, nil
}

func fromDomainProperties(properties []*domain.SchemaProperty) []*MongoSchemaProperty {
	var mongoProperties []*MongoSchemaProperty
	if properties != nil && len(properties) > 0 {
		for _, schemaProperty := range properties {
			mongoProperties = append(mongoProperties, fromDomainProperty(schemaProperty))
		}
	}
	return mongoProperties
}

func fromDomainProperty(property *domain.SchemaProperty) *MongoSchemaProperty {
	if property == nil {
		return nil
	}
	return &MongoSchemaProperty{
		Name:       property.Name,
		Type:       property.Type,
		Min:        property.Min,
		Max:        property.Max,
		Rate:       property.Rate,
		Element:    fromDomainProperty(property.Element),
		RangeSize:  property.RangeSize,
		Properties: fromDomainProperties(property.Properties),
	}
}
