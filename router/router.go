package router

import (
	"database/sql"
	"food-api/api/handler"
	"food-api/api/repository"
	"food-api/api/service"

	"github.com/gin-gonic/gin"
)

func Api(r *gin.Engine, db *sql.DB) {
	repo := repository.NewFoodRepository(db)
	serv := service.NewFoodService(repo)
	handl := handler.NewFoodHandler(serv)

	food := r.Group("food")
	{
		food.POST("/", handl.Create)      //fungsi create
		food.GET("/", handl.GetAll)       // fungso get all
		food.DELETE("/:id", handl.Delete) //fungsi delete
		food.GET("/:id", handl.GetById)   //fungsi get by id
		food.PUT("/:id", handl.Update)    //fungsi update
	}
}
