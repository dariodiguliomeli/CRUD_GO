package application

import products "CRUD_GO/products/core/domain"

type GetAllProductsHandler struct {
	Products products.Products
}

func (handler *GetAllProductsHandler) Exec() (products []products.Product) {
	products = handler.Products.GetAll()
	return products
}
