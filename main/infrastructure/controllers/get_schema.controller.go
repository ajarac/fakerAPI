package controllers

import (
	"fakerAPI/main/application/in"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetSchemaController struct {
	useCase *in.GetSchemaUseCase
}

func NewGetSchemaController(useCase *in.GetSchemaUseCase) *GetSchemaController {
	return &GetSchemaController{useCase: useCase}
}

func (c *GetSchemaController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	schema, exists, err := c.useCase.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !exists {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, schema)
}
