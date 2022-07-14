package controllers

import (
	"fakerAPI/main/application/in"
	"fakerAPI/main/infrastructure/controllers/mapper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProveSchemaController struct {
	useCase *in.ProveSchemaUseCase
}

func (c *ProveSchemaController) Prove(ctx *gin.Context) {
	var schema mapper.JsonSchema
	err := ctx.BindJSON(&schema)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	props, err := mapper.BuildProperties(schema.Properties)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, err := c.useCase.Prove(props)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, value)
}

func NewProveSchemaController(useCase *in.ProveSchemaUseCase) *ProveSchemaController {
	return &ProveSchemaController{useCase: useCase}
}
