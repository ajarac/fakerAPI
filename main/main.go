package main

import (
	"fakerAPI/main/application/in"
	"fakerAPI/main/infrastructure/controllers"
	"fakerAPI/main/infrastructure/provider/faker"
	"fakerAPI/main/infrastructure/storage"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	valueProvider := faker.NewFakerValueProvider()
	schemaStorage := storage.NewInMemorySchemaStorage()
	createSchema := in.NewCreateSchemaUseCase(schemaStorage)
	getValue := in.NewGetValueUseCase(schemaStorage, valueProvider)
	getListValues := in.NewGetListValueUseCase(schemaStorage, valueProvider)
	getSchemas := in.NewGetSchemasUseCase(schemaStorage)

	createSchemaController := controllers.NewCreateSchemaController(createSchema)
	getValueController := controllers.NewGetValueController(getValue)
	getListValuesController := controllers.NewGetListController(getListValues)
	getSchemasControllers := controllers.NewGetSchemasController(getSchemas)
	r.POST("/schemas", createSchemaController.Create)
	r.GET("/schemas", getSchemasControllers.GetAll)
	r.GET("/values/:name", getValueController.ByName)
	r.GET("/values/:name/list", getListValuesController.ByName)
	err := r.Run()
	if err != nil {
		log.Panicln(err)
	}
}
