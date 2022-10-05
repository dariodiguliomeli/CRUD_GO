package main

import (
	"CRUD_GO/products/api"
	"CRUD_GO/products/core/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	productsRepository := products.InMemoryProducts{}
	router := gin.Default()
	api.Init(router, &productsRepository)
	router.GET("/status", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "ok",
		})
	})
	err := router.Run()
	if err != nil {
		panic(err)
	}
}
