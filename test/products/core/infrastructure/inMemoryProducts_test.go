package infrastructure

import (
	infrastructure "CRUD_GO/src/products/core/infrastructure"
	productBuilder "CRUD_GO/test/products/core/domain"
	"testing"
)

func TestNextId(t *testing.T) {
	got := repository.NextId()

	Assert(t, got, 1)
}

func TestGetById(t *testing.T) {
	repository.Add(product)

	got, _ := repository.GetById(product.Id)

	Assert(t, got, product)
}

func TestGetAll(t *testing.T) {
	setupSuite()

	repository.Add(product)
	repository.Add(product)

	products := repository.GetAll()

	Assert(t, len(products), 2)
}

func TestAdd(t *testing.T) {
	setupSuite()
	newProduct := builder.WithId(1).Build()

	repository.Add(newProduct)

	got, _ := repository.GetById(1)

	Assert(t, got, newProduct)
}

var repository = infrastructure.InMemoryProducts{}
var builder = productBuilder.NewProduct()
var product = builder.Build()

func setupSuite() {
	repository = infrastructure.InMemoryProducts{}
}
