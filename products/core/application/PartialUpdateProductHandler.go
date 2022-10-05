package application

import (
	products "CRUD_GO/products/core/domain"
	"errors"
	"time"
)

type PartialUpdateProductHandler struct {
	Products products.Products
}

func (handler *PartialUpdateProductHandler) Exec(id int, name string, description string, price float64) (updatedId int, err error) {
	product, err := handler.Products.GetById(id)
	if err != nil {
		return
	}
	if len(name) <= 0 && len(description) <= 0 && price <= 0 {
		err = errors.New("some field should be update")
		return
	}
	if len(name) > 0 {
		product.Name = name
	}
	if len(description) > 0 {
		product.Description = description
	}
	if price > 0 {
		product.Price = price
	}
	product.UpdatedAt = time.Now()
	updatedId, err = handler.Products.Update(product)
	if err != nil {
		return
	}
	return
}
