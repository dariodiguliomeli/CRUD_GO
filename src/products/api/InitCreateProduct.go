package api

import (
	"CRUD_GO/src/products/core/application"
	"CRUD_GO/src/products/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

func initCreateProductHandler(products products.ProductsPersister) func(context *gin.Context) {
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
