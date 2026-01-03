package orders

type Order struct {
	ProductID int
	Quantity  int
	Status    string
}

func CreateOrder(productID int, quantity int) Order {
	return Order{
		ProductID: productID,
		Quantity:  quantity,
		Status:    "CREATED",
	}
}
