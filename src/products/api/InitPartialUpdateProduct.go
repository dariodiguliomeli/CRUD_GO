package api

import (
	"CRUD_GO/src/products/core/application"
	"CRUD_GO/src/products/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PartialUpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func initPartialUpdateProductHandler(products products.ProductsPersister) func(context *gin.Context) {
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
