package domain

import (
	products "CRUD_GO/src/products/core/domain"
	"time"
)

type productBuilder struct {
	Id          int
	Name        string
	Description string
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductBuilder interface {
	WithId(id int) ProductBuilder
	WithName(name string) ProductBuilder
	WithDescription(description string) ProductBuilder
	WithPrice(price float64) ProductBuilder
	WithCreatedAt(createdAt time.Time) ProductBuilder
	WithUpdatedAt(updatedAt time.Time) ProductBuilder
	Build() products.Product
}

func (builder *productBuilder) WithId(id int) ProductBuilder {
	builder.Id = id
	return builder
}

func (builder *productBuilder) WithName(name string) ProductBuilder {
	builder.Name = name
	return builder
}

func (builder *productBuilder) WithDescription(description string) ProductBuilder {
	builder.Description = description
	return builder
}

func (builder *productBuilder) WithPrice(price float64) ProductBuilder {
	builder.Price = price
	return builder
}

func (builder *productBuilder) WithCreatedAt(createdAt time.Time) ProductBuilder {
	builder.CreatedAt = createdAt
	return builder
}

func (builder *productBuilder) WithUpdatedAt(updatedAt time.Time) ProductBuilder {
	builder.UpdatedAt = updatedAt
	return builder
}

func (builder *productBuilder) Build() (product products.Product) {
	newProduct := products.Product{
		Id:          builder.Id,
		Name:        builder.Name,
		Description: builder.Description,
		Price:       builder.Price,
		CreatedAt:   builder.CreatedAt,
		UpdatedAt:   builder.UpdatedAt,
	}
	fillDefaultFields(builder)
	return newProduct
}

func fillDefaultFields(builder *productBuilder) {
	if builder.Id == 0 {
		builder.Id = 1
	}
	if builder.Name == "" {
		builder.Name = "Product name"
	}
	if builder.Description == "" {
		builder.Description = "Description"
	}
	if builder.Price == 0.0 {
		builder.Price = 23.4
	}
	if builder.CreatedAt.IsZero() {
		builder.CreatedAt = time.Now()
	}
	if builder.UpdatedAt.IsZero() {
		builder.UpdatedAt = time.Now()
	}
}

func New() ProductBuilder {
	return &productBuilder{}
}
