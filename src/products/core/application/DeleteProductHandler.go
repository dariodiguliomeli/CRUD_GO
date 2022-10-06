package application

import (
	"CRUD_GO/src/products/core/domain"
)

type DeleteProductHandler struct {
	Products products.ProductsPersister
}

func (handler *DeleteProductHandler) Exec(id int) (deletedId int, err error) {
	return handler.Products.Delete(id)
}
