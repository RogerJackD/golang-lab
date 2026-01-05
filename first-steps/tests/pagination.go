package pagination

func Offset(page, limit int) int {
	if page <= 1 {
		return 0
	}
	return (page - 1) * limit
}
