package main

import (
	"fakerAPI/main/config"
	"fakerAPI/main/infrastructure"
	"fakerAPI/main/infrastructure/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	env := config.GetEnvironment()
	useCases := infrastructure.NewUseCases(env)

	createSchemaController := controllers.NewCreateSchemaController(useCases.CreateSchema)
	getValueController := controllers.NewGetValueController(useCases.GetValue)
	getListValuesController := controllers.NewGetListController(useCases.GetListValue)
	getSchemasControllers := controllers.NewGetSchemasController(useCases.GetSchemas)
	r.POST("/schemas", createSchemaController.Create)
	r.GET("/schemas", getSchemasControllers.GetAll)
	r.GET("/values/:name", getValueController.ByName)
	r.GET("/values/:name/list", getListValuesController.ByName)

	err := r.Run()
	if err != nil {
		log.Panicln(err)
	}
}
