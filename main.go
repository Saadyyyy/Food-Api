package main

import (
	"food-api/config"
	"food-api/router"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	db, _ := config.Connection()

	router.Api(r, db)
	r.Run()
}
