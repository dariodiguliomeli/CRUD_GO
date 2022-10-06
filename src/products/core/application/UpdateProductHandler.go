package application

import (
	"CRUD_GO/src/products/core/domain"
	"time"
)

type UpdateProductHandler struct {
	Products products.ProductsPersister
}

func (handler *UpdateProductHandler) Exec(id int, name string, description string, price float64) (updatedId int, err error) {
	product, err := handler.Products.GetById(id)
	if err != nil {
		return
	}
	product.Name = name
	product.Description = description
	product.Price = price
	product.UpdatedAt = time.Now()
	return handler.Products.Update(product)
}
