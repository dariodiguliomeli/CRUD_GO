package api

import (
	"CRUD_GO/products/core/application"
	"CRUD_GO/products/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(router *gin.Engine, products products.Products) {
	productsRouter := router.Group("/products")
	{
		productsRouter.POST("/", initCreateProductHandler(products))
	}
}

type CreateProductRequest struct {
	name        string
	description string
	price       float64
}

func initCreateProductHandler(products products.Products) func(context *gin.Context) {
	return func(context *gin.Context) {
		handler := application.CreateProductHandler{Products: products}
		var request CreateProductRequest
		if err := context.BindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		id := handler.Exec(request.name, request.description, request.price)
		context.JSON(http.StatusCreated, gin.H{"Id": id})
		return
	}
}
