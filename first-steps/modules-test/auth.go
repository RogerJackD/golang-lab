package auth

func Login(email string, password string) bool {
	return email == "admin@test.com" && password == "123456"
}
