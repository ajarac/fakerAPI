package controllers

import (
	"fakerAPI/main/application/in"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteSchemaController struct {
	useCase *in.DeleteSchemaUseCase
}

func NewDeleteSchemaController(useCase *in.DeleteSchemaUseCase) *DeleteSchemaController {
	return &DeleteSchemaController{useCase: useCase}
}

func (c *DeleteSchemaController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.useCase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusAccepted)
}
