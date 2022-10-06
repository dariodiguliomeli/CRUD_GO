package api

import (
	application2 "CRUD_GO/src/products/core/application"
	"CRUD_GO/src/products/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type ProductResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func initGetProductsHandler(products products.ProductsPersister) func(context *gin.Context) {
	return func(context *gin.Context) {
		handler := application2.GetAllProductsHandler{Products: products}
		productsFound := handler.Exec()
		var productsToReturn []ProductResponse
		for _, product := range productsFound {
			productsToReturn = append(productsToReturn, toProductResponse(product))
		}
		context.JSON(http.StatusOK, gin.H{"Products": productsFound})
	}
}

func initGetProductByIdHandler(products products.ProductsPersister) func(context *gin.Context) {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": "id param not found"})
			return
		}
		handler := application2.GetProductByIdHandler{Products: products}
		product, err := handler.Exec(id)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"Product": toProductResponse(product)})
	}
}

func toProductResponse(product products.Product) (productResponse ProductResponse) {
	return ProductResponse{
		ID:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
