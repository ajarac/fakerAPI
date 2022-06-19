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
	ctx.JSON(http.StatusOK, c.useCase.Get())
}

func NewGetSchemasController(useCase *in.GetSchemasUseCase) *GetSchemasController {
	return &GetSchemasController{useCase: useCase}
}
