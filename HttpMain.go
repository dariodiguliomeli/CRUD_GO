package main

import (
	"CRUD_GO/src/products/api"
	"CRUD_GO/src/products/core/infrastructure"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	loadEnv()
	productsRepository := products.InMemoryProductsPersister{}
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

func Init(productsRepository products.InMemoryProductsPersister) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(authMiddleware())
	api.Init(router, &productsRepository)
	statusHandler(router)
	return router
}

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userToken := ctx.Request.Header.Get("Authorization")
		apiToken := fmt.Sprintf("%s", os.Getenv("TOKEN"))
		if userToken != apiToken {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
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
