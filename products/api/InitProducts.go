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
		productsRouter.GET("/", initGetProductHandler(products))
	}
}

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

func initCreateProductHandler(products products.Products) func(context *gin.Context) {
	return func(context *gin.Context) {
		handler := application.CreateProductHandler{Products: products}
		var request CreateProductRequest
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		id := handler.Exec(request.Name, request.Description, request.Price)
		context.JSON(http.StatusCreated, gin.H{"Id": id})
		return
	}
}

func initGetProductHandler(products products.Products) func(context *gin.Context) {
	return func(context *gin.Context) {
		handler := application.GetAllProductsHandler{Products: products}
		context.JSON(http.StatusOK, gin.H{"Products": handler.Exec()})
	}
}
