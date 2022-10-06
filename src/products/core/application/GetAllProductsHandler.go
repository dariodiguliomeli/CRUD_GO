package application

import (
	products2 "CRUD_GO/src/products/core/domain"
)

type GetAllProductsHandler struct {
	Products products2.ProductsPersister
}

func (handler *GetAllProductsHandler) Exec() (products []products2.Product) {
	return handler.Products.GetAll()
}
