package auth

type TokenStorage interface {
	StoreToken(UserId, TokenString) error
	RevokeToken(UserId) error
}
