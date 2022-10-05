package application

import (
	products "CRUD_GO/products/core/domain"
	"time"
)

type UpdateProductHandler struct {
	Products products.Products
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
