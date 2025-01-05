package postgres

type User struct {
	UserId       string
	Username     string
	PasswordHash string
	Role         string
}
