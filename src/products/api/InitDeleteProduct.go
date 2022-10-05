package api

import (
	"CRUD_GO/src/products/core/application"
	"CRUD_GO/src/products/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
