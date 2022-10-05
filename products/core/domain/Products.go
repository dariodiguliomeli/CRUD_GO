package products

type Products interface {
	NextId() (id int)
	GetById(id int) (product Product)
	GetAll() (products []Product)
	Delete(id int) (idDeleted int)
	Add(product Product) (id int)
	Update(id int, name string, description string, price float64) (idUpdated int)
}
