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
		food.GET("/", handl.Create)
		// fmt.Println(food.GET("/", handl.Create))
	}
}
