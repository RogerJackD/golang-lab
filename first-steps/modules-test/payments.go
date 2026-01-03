package payments

type Payment struct {
	Amount float64
	Method string
	Status string
}

func ProcessPayment(amount float64, method string) Payment {
	return Payment{
		Amount: amount,
		Method: method,
		Status: "PAID",
	}
}
