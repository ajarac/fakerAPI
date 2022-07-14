package infrastructure

import (
	"fakerAPI/main/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func BuildControllersV1(group *gin.RouterGroup, useCases *UseCases) {
	createSchemaController := controllers.NewCreateSchemaController(useCases.CreateSchema)
	getValueController := controllers.NewGetValueController(useCases.GetValue)
	getListValuesController := controllers.NewGetListController(useCases.GetListValue)
	getSchemasControllers := controllers.NewGetSchemasController(useCases.GetSchemas)
	deleteSchemaController := controllers.NewDeleteSchemaController(useCases.DeleteSchema)
	getSchemaController := controllers.NewGetSchemaController(useCases.GetSchema)
	proveSchemaController := controllers.NewProveSchemaController(useCases.ProveSchema)

	group.POST("/schemas", createSchemaController.Create)
	group.POST("/schemas/prove", proveSchemaController.Prove)
	group.GET("/schemas", getSchemasControllers.GetAll)
	group.DELETE("/schemas/:id", deleteSchemaController.Delete)
	group.GET("/schemas/:id", getSchemaController.GetById)
	group.GET("/values/:id", getValueController.ById)
	group.GET("/values/:id/list", getListValuesController.ById)
}
