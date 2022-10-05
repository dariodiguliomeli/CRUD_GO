package api

import (
	"CRUD_GO/products/core/application"
	"CRUD_GO/products/core/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Init(router *gin.Engine, products products.Products) {
	productsRouter := router.Group("/products")
	{
		productsRouter.POST("/", initCreateProductHandler(products))
		productsRouter.GET("/", initGetProductHandler(products))
		productsRouter.GET("/:id", initGetProductByIdHandler(products))
		productsRouter.PATCH("/:id", initPartialUpdateProductHandler(products))
		productsRouter.DELETE("/:id", initDeleteProductByIdHandler(products))
		productsRouter.PUT("/:id", initUpdateProductHandler(products))
	}
}

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

type ProductResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type PartialUpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type UpdateProductRequest struct {
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

func initGetProductByIdHandler(products products.Products) func(context *gin.Context) {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": "id param not found"})
			return
		}
		handler := application.GetProductByIdHandler{Products: products}
		product, err := handler.Exec(id)
		fmt.Println(product)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"Product": product})
	}
}

func initPartialUpdateProductHandler(products products.Products) func(context *gin.Context) {
	return func(context *gin.Context) {
		var request PartialUpdateProductRequest
		if err := context.BindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		handler := application.PartialUpdateProductHandler{Products: products}
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": "id param not found"})
			return
		}
		id, err = handler.Exec(id, request.Name, request.Description, request.Price)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"Id": id})
	}
}

func initDeleteProductByIdHandler(products products.Products) func(context *gin.Context) {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": "id param not found"})
			return
		}
		handler := application.DeleteProductHandler{Products: products}
		deletedId, err := handler.Exec(id)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		}
		context.JSON(http.StatusOK, gin.H{"Id": deletedId})
	}
}

func initUpdateProductHandler(products products.Products) func(context *gin.Context) {
	return func(context *gin.Context) {
		var request UpdateProductRequest
		if err := context.BindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		handler := application.UpdateProductHandler{Products: products}
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": "id param not found"})
			return
		}
		id, err = handler.Exec(id, request.Name, request.Description, request.Price)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"Id": id})
	}
}
