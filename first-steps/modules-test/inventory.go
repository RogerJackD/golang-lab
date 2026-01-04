package inventory

type Item struct {
	ProductID int
	Stock     int
}

func CheckStock(productID int) Item {
	return Item{
		ProductID: productID,
		Stock:     100,
	}
}
