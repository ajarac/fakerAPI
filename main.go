package main

import (
	"fakerAPI/main/config"
	"fakerAPI/main/infrastructure"
	"fakerAPI/main/infrastructure/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	env := config.GetEnvironment()
	useCases := infrastructure.NewUseCases(env)

	createSchemaController := controllers.NewCreateSchemaController(useCases.CreateSchema)
	getValueController := controllers.NewGetValueController(useCases.GetValue)
	getListValuesController := controllers.NewGetListController(useCases.GetListValue)
	getSchemasControllers := controllers.NewGetSchemasController(useCases.GetSchemas)
	deleteSchemaController := controllers.NewDeleteSchemaController(useCases.DeleteSchema)
	r.Use(func(context *gin.Context) {
		rapidAPISecret := context.GetHeader("X-RapidAPI-Proxy-Secret")
		log.Printf("RAPID SECRET:: %s", rapidAPISecret)
		if rapidAPISecret != env.RapidAPIKey {
			context.AbortWithStatus(http.StatusUnauthorized)
		}

	})
	r.GET("/health", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})
	r.POST("/schemas", createSchemaController.Create)
	r.GET("/schemas", getSchemasControllers.GetAll)
	r.DELETE("/schemas/:id", deleteSchemaController.Delete)
	r.GET("/values/:id", getValueController.ById)
	r.GET("/values/:id/list", getListValuesController.ById)

	err := r.Run(":" + env.Port)
	if err != nil {
		log.Panicln(err)
	}
}
