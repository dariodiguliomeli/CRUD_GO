package application

import (
	"CRUD_GO/src/products/core/application"
	infrastructure "CRUD_GO/src/products/core/infrastructure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProductHandler(t *testing.T) {
	createdId := handler.Exec("Some name", "Some description", 23.45)

	productSaved, _ := repository.GetById(createdId)
	assert.Equal(t, productSaved.Id, createdId)
	assert.Equal(t, productSaved.Name, "Some name")
	assert.Equal(t, productSaved.Description, "Some description")
	assert.Equal(t, productSaved.Price, 23.45)
}

func TestCreateProductHandlerSequentialId(t *testing.T) {
	setupSuite()
	firstId := handler.Exec("Some name", "Some description", 23.45)
	secondId := handler.Exec("Some name", "Some description", 23.45)

	assert.Equal(t, firstId, 1)
	assert.Equal(t, secondId, 2)
}

func setupSuite() {
	repository = infrastructure.InMemoryProductsPersister{}
}

var repository = infrastructure.InMemoryProductsPersister{}
var handler = application.CreateProductHandler{Products: &repository}
