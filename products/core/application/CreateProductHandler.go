package application

import (
	"CRUD_GO/products/core/domain"
	"errors"
	"time"
)

type CreateProductHandler struct {
	Products products.Products
}

func (handler *CreateProductHandler) Exec(name string, description string, price float64) (id int) {
	err := validate(name, description, price)
	if err != nil {
		panic(err)
	}
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

func validate(name string, description string, price float64) (err error) {
	err = validateName(name)
	err = validateDescription(description)
	err = validatePrice(price)
	return err
}

func validatePrice(price float64) error {
	if price <= 0 {
		return errors.New("price is empty")
	}
	return nil
}

func validateDescription(description string) error {
	if len(description) <= 0 {
		return errors.New("description is empty")
	}
	return nil
}

func validateName(name string) error {
	if len(name) <= 0 {
		return errors.New("name is empty")
	}
	return nil
}
