package infrastructure

import (
	"fakerAPI/main/application/in"
	"fakerAPI/main/config"
	"fakerAPI/main/infrastructure/provider/faker"
	"fakerAPI/main/infrastructure/storage"
)

type UseCases struct {
	CreateSchema *in.CreateSchemaUseCase
	GetValue     *in.GetValueUseCase
	GetListValue *in.GetListValueUseCase
	GetSchemas   *in.GetSchemasUseCase
}

func NewUseCases(env *config.Environment) *UseCases {
	dbClient := GetDatabase(env)
	valueProvider := faker.NewFakerValueProvider()
	schemaStorage := storage.NewMongoDBSchemaStorage(dbClient)
	return &UseCases{
		CreateSchema: in.NewCreateSchemaUseCase(schemaStorage),
		GetValue:     in.NewGetValueUseCase(schemaStorage, valueProvider),
		GetListValue: in.NewGetListValueUseCase(schemaStorage, valueProvider),
		GetSchemas:   in.NewGetSchemasUseCase(schemaStorage),
	}
}
