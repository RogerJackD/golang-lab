package reviews

type Review struct {
	ProductID int
	Rating    int
	Comment   string
}

func CreateReview(productID int, rating int, comment string) Review {
	return Review{
		ProductID: productID,
		Rating:    rating,
		Comment:   comment,
	}
}
