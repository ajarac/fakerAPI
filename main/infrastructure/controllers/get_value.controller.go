package controllers

import (
	"errors"
	"fakerAPI/main/application/in"
	"fakerAPI/main/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetValueController struct {
	useCase *in.GetValueUseCase
}

func (c *GetValueController) ById(ctx *gin.Context) {
	id := ctx.Param("id")
	value, err := c.useCase.GetById(id)
	var schemaNotFound *domain.SchemaNotFound
	if errors.As(err, &schemaNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, value)
}

func NewGetValueController(useCase *in.GetValueUseCase) *GetValueController {
	return &GetValueController{useCase: useCase}
}
