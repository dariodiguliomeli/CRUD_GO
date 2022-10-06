package products

import (
	"CRUD_GO/src/products/core/domain"
	"errors"
)

var nextId = 1

type InMemoryProductsPersister struct {
	Products []products.Product
}

func (repository *InMemoryProductsPersister) NextId() (id int) {
	id = nextId
	nextId++
	return
}

func (repository *InMemoryProductsPersister) Add(product products.Product) (id int) {
	repository.Products = append(repository.Products, product)
	return product.Id
}

func (repository *InMemoryProductsPersister) GetById(id int) (product products.Product, err error) {
	for _, item := range repository.Products {
		if item.Id == id {
			return item, nil
		}
	}
	err = errors.New("product not found")
	return
}

func (repository *InMemoryProductsPersister) GetAll() (products []products.Product) {
	return repository.Products
}

func (repository *InMemoryProductsPersister) Delete(id int) (idDeleted int, err error) {
	for i, item := range repository.Products {
		if item.Id == id {
			repository.Products = append(repository.Products[:i], repository.Products[i+1:]...)
			return id, nil
		}
	}
	err = errors.New("product not found")
	return
}

func (repository *InMemoryProductsPersister) Update(product products.Product) (updatedId int, err error) {
	for i, item := range repository.Products {
		if item.Id == product.Id {
			repository.Products[i] = product
			return product.Id, nil
		}
	}
	err = errors.New("product not found")
	return
}
