package application

import products "CRUD_GO/products/core/domain"

type GetProductByIdHandler struct {
	Products products.Products
}

func (handler *GetProductByIdHandler) Exec(id int) (product products.Product, err error) {
	product, err = handler.Products.GetById(id)
	return
}
