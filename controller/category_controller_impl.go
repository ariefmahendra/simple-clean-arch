package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api-clean-arch/helper"
	"simple-rest-api-clean-arch/model/dto"
	"simple-rest-api-clean-arch/service"
	"strconv"
)

type CategoryControllerImpl struct {
	categoryService service.CategoryService
}

func (controller *CategoryControllerImpl) Create(ctx *gin.Context) {
	createdCategory := dto.CategoryCreateRequest{}
	err := ctx.BindJSON(&createdCategory)
	helper.PanicIfError(err)

	categoryResponse := controller.categoryService.Create(ctx, createdCategory)

	webResponse := dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryControllerImpl) Update(ctx *gin.Context) {
	updatedCategory := dto.CategoryUpdateRequest{}
	err := ctx.ShouldBindJSON(&updatedCategory)
	helper.PanicIfError(err)

	categoryId := ctx.Param("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	updatedCategory.Id = id

	categoryResponse := controller.categoryService.Update(ctx, updatedCategory)

	webResponse := dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryControllerImpl) Delete(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.categoryService.Delete(ctx, id)

	webResponse := dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryControllerImpl) FindByID(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.categoryService.FindByID(ctx, id)

	webResponse := dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(ctx *gin.Context) {
	categoryResponse := controller.categoryService.FindAll(ctx)

	webResponse := dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
