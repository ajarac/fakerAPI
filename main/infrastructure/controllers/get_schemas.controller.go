package controllers

import (
	"fakerAPI/main/application/in"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetSchemasController struct {
	useCase *in.GetSchemasUseCase
}

func (c *GetSchemasController) GetAll(ctx *gin.Context) {
	schemas, err := c.useCase.Get(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, schemas)
}

func NewGetSchemasController(useCase *in.GetSchemasUseCase) *GetSchemasController {
	return &GetSchemasController{useCase: useCase}
}
