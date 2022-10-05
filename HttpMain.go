package main

import (
	"CRUD_GO/products/api"
	"CRUD_GO/products/core/infrastructure"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	loadEnv()
	productsRepository := products.InMemoryProducts{}
	router := Init(productsRepository)
	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	run(router, port)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func Init(productsRepository products.InMemoryProducts) *gin.Engine {
	router := gin.Default()
	api.Init(router, &productsRepository)
	statusHandler(router)
	return router
}

func statusHandler(router *gin.Engine) gin.IRoutes {
	return router.GET("/status", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "ok",
		})
	})
}

func run(router *gin.Engine, port string) {
	err := router.Run(port)
	if err != nil {
		panic(err)
	}
}
