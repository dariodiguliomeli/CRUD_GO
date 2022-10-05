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
	err := router.Run()
	if err != nil {
		panic(err)
	}
}
