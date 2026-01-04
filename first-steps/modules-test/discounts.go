package discounts

func ApplyDiscount(price float64, percent float64) float64 {
	return price - (price * percent / 100)
}
