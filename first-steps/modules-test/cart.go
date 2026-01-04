package cart

type CartItem struct {
	ProductID int
	Quantity  int
}

func AddToCart(productID int, quantity int) CartItem {
	return CartItem{
		ProductID: productID,
		Quantity:  quantity,
	}
}
