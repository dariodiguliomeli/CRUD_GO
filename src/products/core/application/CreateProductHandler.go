package application

import (
	"CRUD_GO/src/products/core/domain"
	"time"
)

type CreateProductHandler struct {
	Products products.Products
}

func (handler *CreateProductHandler) Exec(name string, description string, price float64) (id int) {
	nextId := handler.Products.NextId()
	newProduct := products.Product{
		Id:          nextId,
		Name:        name,
		Description: description,
		Price:       price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	handler.Products.Add(newProduct)
	return nextId
}
