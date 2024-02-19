package handler

import (
	"food-api/api/service"
	"food-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	service service.FoodService
}

func NewFoodHandler(service service.FoodService) *FoodHandler {
	return &FoodHandler{service: service}
}

func (fh *FoodHandler) Create(ctx *gin.Context) {
	food := models.Food{}
	if err := ctx.ShouldBind(&food.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Error",
			"Data":    err,
		})
		return
	}
	data, err := fh.service.Create(food)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Error",
			"Data":    err.Error(),
			"err":     "data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Status Ok",
		"Data":    data,
	})
}

// GetAll implements FoodService.
func (fh *FoodHandler) GetAll(ctx *gin.Context) {
	food := []models.Food{}
	if err := ctx.ShouldBind(&food); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Error",
			"Data":    err,
		})
		return
	}
	data, err := fh.service.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Error",
			"Data":    err.Error(),
			"err":     "data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Status Ok",
		"Data":    data,
	})
}

// Delete implements FoodService.
func (fh *FoodHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
			"error":   err.Error(),
		})
		return
	}
	data, err := fh.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "status internal server error",
			"Data":    err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "status oke",
		"data":    data,
	})

}

// GetById implements FoodService.
func (fh *FoodHandler) GetById() {
	panic("unimplemented")
}

// Update implements FoodService.
func (fh *FoodHandler) Update() {
	panic("unimplemented")
}
