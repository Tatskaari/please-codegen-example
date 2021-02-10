package auth

func Authenticate(username, password string) bool {
	return password == "hunter2"
}