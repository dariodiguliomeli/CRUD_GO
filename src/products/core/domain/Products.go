package products

type Products interface {
	NextId() (id int)
	GetById(id int) (product Product, err error)
	GetAll() (products []Product)
	Delete(id int) (idDeleted int, err error)
	Add(product Product) (id int)
	Update(product Product) (updatedId int, err error)
}
