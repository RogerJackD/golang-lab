package products

type Product struct {
	ID    int
	Name  string
	Price float64
}

func GetProductByID(id int) Product {
	return Product{
		ID:    id,
		Name:  "Laptop",
		Price: 1200.00,
	}
}
