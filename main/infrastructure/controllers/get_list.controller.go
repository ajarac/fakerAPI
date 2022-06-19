package controllers

import (
	"errors"
	"fakerAPI/main/application/in"
	"fakerAPI/main/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetListController struct {
	useCase *in.GetListValueUseCase
}

func NewGetListController(useCase *in.GetListValueUseCase) *GetListController {
	return &GetListController{useCase: useCase}
}

func (c *GetListController) ByName(ctx *gin.Context) {
	name := ctx.Param("name")
	limitString := ctx.DefaultQuery("limit", "25")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	list, err := c.useCase.GetByNameAndLimit(name, limit)
	var schemaNotFound *domain.SchemaNotFound
	if errors.As(err, &schemaNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, list)
}
