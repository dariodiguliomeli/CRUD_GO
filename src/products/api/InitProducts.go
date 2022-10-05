package api

import (
	products "CRUD_GO/src/products/core/domain"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine, products products.Products) {
	productsRouter := router.Group("/products")
	{
		productsRouter.POST("/", initCreateProductHandler(products))
		productsRouter.GET("/", initGetProductsHandler(products))
		productsRouter.GET("/:id", initGetProductByIdHandler(products))
		productsRouter.PATCH("/:id", initPartialUpdateProductHandler(products))
		productsRouter.DELETE("/:id", initDeleteProductByIdHandler(products))
		productsRouter.PUT("/:id", initUpdateProductHandler(products))
	}
}
