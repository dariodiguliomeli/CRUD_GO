package application

import products "CRUD_GO/products/core/domain"

type DeleteProductHandler struct {
	Products products.Products
}

func (handler *DeleteProductHandler) Exec(id int) (deletedId int, err error) {
	deletedId, err = handler.Products.Delete(id)
	if err != nil {
		return
	}
	return
}
