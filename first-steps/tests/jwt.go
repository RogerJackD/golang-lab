package jwt

func GenerateToken(userID string) string {
	return "token-for-" + userID
}
