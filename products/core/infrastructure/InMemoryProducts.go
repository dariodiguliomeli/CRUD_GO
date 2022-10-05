package products

import products "CRUD_GO/products/core/domain"

var nextId = 1

type InMemoryProducts struct {
	Products []products.Product
}

func (repository *InMemoryProducts) NextId() (id int) {
	id = nextId
	nextId++
	return
}

func (repository *InMemoryProducts) Add(product products.Product) (id int) {
	repository.Products = append(repository.Products, product)
	return product.Id
}

func (repository *InMemoryProducts) GetById(id int) (product products.Product) {
	//TODO implement me
	panic("implement me")
}

func (repository *InMemoryProducts) GetAll() (products []products.Product) {
	return repository.Products
}

func (repository *InMemoryProducts) Delete(id int) (idDeleted int) {
	//TODO implement me
	panic("implement me")
}

func (repository *InMemoryProducts) Update(id int, name string, description string, price float64) (idUpdated int) {
	//TODO implement me
	panic("implement me")
}
