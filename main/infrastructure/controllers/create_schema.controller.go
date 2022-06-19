package controllers

import (
	"fakerAPI/main/application/in"
	"fakerAPI/main/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateSchemaController struct {
	useCase *in.CreateSchemaUseCase
}

func (c *CreateSchemaController) Create(ctx *gin.Context) {
	var schema domain.Schema
	err := ctx.BindJSON(&schema)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.useCase.Create(&schema)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}

func NewCreateSchemaController(useCase *in.CreateSchemaUseCase) *CreateSchemaController {
	return &CreateSchemaController{useCase: useCase}
}
