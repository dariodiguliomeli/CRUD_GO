package infrastructure

import (
	products "CRUD_GO/src/products/core/infrastructure"
	"testing"
)

func TestNextId(t *testing.T) {
	got := repository.NextId()
	want := 1

	assert(t, got, want)
}

func TestGetById(t *testing.T) {

}

var repository = products.InMemoryProducts{}

func assert(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("NextId() = %v, want %v", got, want)
	}
}
