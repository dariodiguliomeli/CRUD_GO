package infrastructure

import (
	domain "CRUD_GO/src/products/core/domain"
	infrastructure "CRUD_GO/src/products/core/infrastructure"
	"testing"
	"time"
)

func TestNextId(t *testing.T) {
	got := repository.NextId()

	assert(t, got, 1)
}

func TestGetById(t *testing.T) {
	repository.Add(product)

	got, _ := repository.GetById(product.Id)

	assert(t, got, product)

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

func assert(t *testing.T, got any, want any) {
	if got != want {
		t.Errorf("NextId() = %v, want %v", got, want)
	}
}
