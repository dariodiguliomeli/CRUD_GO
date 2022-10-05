package application

import (
	products2 "CRUD_GO/src/products/core/domain"
)

type GetProductByIdHandler struct {
	Products products2.Products
}

func (handler *GetProductByIdHandler) Exec(id int) (product products2.Product, err error) {
	return handler.Products.GetById(id)
}
