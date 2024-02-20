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

// handler untuk fungsi create
func (fh *FoodHandler) Create(ctx *gin.Context) {
	//memanggil objek dari foods
	food := models.Food{}
	//akan menambil data dari body foods dari body json yang di minput
	if err := ctx.ShouldBind(&food); err != nil {
		//jika terdapat error akan di handle dan menampilkan error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Error",
			"Data":    err,
		})
		return
	}

	//memanggil fungsi create dari service dan di tampung oleh 2 variable
	data, err := fh.service.Create(&food)
	//jika fungsi create dari service terdapat error akan di handle dan mereturn error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Error",
			"Data":    err.Error(),
			"err":     "data",
		})
		return
	}
	// jika fungsi dari create tidak terdapat error akan di return status oke
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Status Ok",
		"Data":    data,
	})
}

// GetAll implements FoodService.
func (fh *FoodHandler) GetAll(ctx *gin.Context) {
	//memanggil objek dari model.food
	food := []models.Food{}
	//mengambil objek dari table foods
	if err := ctx.ShouldBind(&food); err != nil {
		// jika terdapat error akan di handle dan mereturn error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Error",
			"Data":    err,
		})
		return
	}
	//memanggil fungsi getall dari service dan menampung 2 variable
	data, err := fh.service.GetAll()
	//jika fungsi get all dari service terdapat error akan di handle dan di return error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Error",
			"Data":    err.Error(),
			"err":     "data",
		})
		return
	}
	//jika fungsi get all dari service sudah tidak terdapat error akan di return status oke
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Status Ok",
		"Data":    data,
	})
}

// Delete implements FoodService.
func (fh *FoodHandler) Delete(ctx *gin.Context) {
	//melakukan konversi dari id string ke id integer
	id, err := strconv.Atoi(ctx.Param("id"))
	//jika saat konversi terdapat error akan di handle dan mereturn error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"massage": "Invalid Id",
			"err":     err.Error(),
		})
	}
	//melakukan pengecekan jika id di bawah 0 akan melakukan return error
	if id < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"massage": "Invalid id",
			"err":     err.Error(),
		})
		return
	}
	//memanggil fungsi get by id dari service untuk memeriksa apakah id yang di input ada di database
	_, err = fh.service.GetById(id)
	//jika id yang di input tidak ada di database maka akan di handle dan di return error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Id Not Found",
			"error":   err.Error(), // ubah "Data" menjadi "error"
		})
		return
	}
	// memanggil fungsi delete dari service dan di tampung dengan 1 variable
	err = fh.service.Delete(id)
	// jiks fungsi delete dari service terdapat error makan akan di handle dan di return error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(), // ubah "Data" menjadi "error"
		})
		return
	}

	// jika fungsi delete service tidak terdapat error akan di return status oke
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Status Ok",
		"Data":    "Succesfully delete",
	})

}

// GetById implements FoodService.
func (fh *FoodHandler) GetById(ctx *gin.Context) {
	//mengambil id dari param dan di tampung di idStr
	idStr := ctx.Param("id")
	//mengkonversikan idStr dari string ke Integer
	id, err := strconv.Atoi(idStr)
	//jika saat konversi terdapat error akan di handle dan di return error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
			"error":   err.Error(),
		})
		return
	}
	//memanggil fungsi getbyid dari service dan di tampung oleh dua variable
	data, err := fh.service.GetById(id)
	//jika fungsi get by id terdapat error akan di handle dan di return error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"massage": "Id Not Found",
		})
		return
	}
	//jika fungsi dari get by id tidak terdapat error akan di return status oke
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "status oke",
		"data":    data,
	})
}

// Update implements FoodService.
func (fh *FoodHandler) Update(ctx *gin.Context) {
	// mengambil dan mengubah id
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Invalid id",
			"error":   err.Error(),
		})
		return
	}
	// cek apakah id ada di database
	_, errId := fh.service.GetById(id)
	//jika tidak akan merreturn not found
	if errId != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"massage": "Id Not found",
			// "err":     err.Error(),
		})
		return
	}
	//memanggil model dari food
	food := models.Food{}
	//mengambil data dari body yang di input
	if err := ctx.ShouldBind(&food); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Error",
			"Data":    err,
		})
		return
	}

	// memanggin service update
	data, errUpdate := fh.service.Update(id, food)
	//cek apakah mereturn err jika err internal server err
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Internal server error",
			// "err":     err.Error(),
		})
		return
	}

	//respon jika suses dan menampilkan data yang telah di update
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Succesfully update data",
		"data":    data,
	})
}
