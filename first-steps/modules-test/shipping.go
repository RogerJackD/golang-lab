package shipping

type Shipping struct {
	Address string
	Status  string
}

func CreateShipping(address string) Shipping {
	return Shipping{
		Address: address,
		Status:  "PENDING",
	}
}
