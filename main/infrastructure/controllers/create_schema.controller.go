package controllers

import (
	"fakerAPI/main/application/in"
	"fakerAPI/main/infrastructure/controllers/mapper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateSchemaController struct {
	useCase *in.CreateSchemaUseCase
}

func (c *CreateSchemaController) Create(ctx *gin.Context) {
	var schema mapper.JsonSchema
	err := ctx.BindJSON(&schema)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	schemaCreated, err := c.useCase.Create(ctx, schema.Name, mapper.BuildProperties(schema.Properties))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, schemaCreated)
}

func NewCreateSchemaController(useCase *in.CreateSchemaUseCase) *CreateSchemaController {
	return &CreateSchemaController{useCase: useCase}
}
