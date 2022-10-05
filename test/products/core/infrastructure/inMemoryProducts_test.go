package infrastructure

import (
	domain "CRUD_GO/src/products/core/domain"
	infrastructure "CRUD_GO/src/products/core/infrastructure"
	"testing"
	"time"
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

var repository = infrastructure.InMemoryProducts{}
var product = domain.Product{
	Id:          1,
	Name:        "test",
	Description: "some description",
	Price:       45.23,
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
}

func setupSuite() {
	repository = infrastructure.InMemoryProducts{}
}
