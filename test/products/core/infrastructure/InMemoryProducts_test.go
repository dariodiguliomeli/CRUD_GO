package infrastructure

import (
	infrastructure "CRUD_GO/src/products/core/infrastructure"
	productBuilder "CRUD_GO/test/products/core/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextId(t *testing.T) {
	got := repository.NextId()

	assert.Equal(t, got, 1)
}

func TestGetById(t *testing.T) {
	repository.Add(product)

	got, _ := repository.GetById(product.Id)

	assert.Equal(t, got, product)
}

func TestGetAll(t *testing.T) {
	setupSuite()

	repository.Add(product)
	repository.Add(product)

	products := repository.GetAll()

	assert.Equal(t, len(products), 2)
}

func TestAdd(t *testing.T) {
	setupSuite()
	newProduct := builder.WithId(1).Build()

	repository.Add(newProduct)

	got, _ := repository.GetById(1)

	assert.Equal(t, got, newProduct)
}

func TestDelete(t *testing.T) {
	setupSuite()
	newProduct := builder.WithId(1).Build()
	repository.Add(newProduct)

	id, _ := repository.Delete(newProduct.Id)

	searchProduct, _ := repository.GetById(id)
	assert.Equal(t, searchProduct.Id, 0)
}

func TestUpdate(t *testing.T) {
	setupSuite()
	productToUpdate := builder.WithName("Taza").Build()
	repository.Add(productToUpdate)
	productToUpdate.Name = "Plato"

	updatedId, _ := repository.Update(productToUpdate)

	updatedProduct, _ := repository.GetById(updatedId)
	assert.Equal(t, updatedProduct.Name, "Plato")
}

var repositoryType = infrastructure.InMemoryProductsPersister{}
var repository = repositoryType.New()
var builder = productBuilder.New()
var product = builder.Build()

func setupSuite() {
	repository = infrastructure.InMemoryProductsPersister{}
}
