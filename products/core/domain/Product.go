package products

import "time"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
