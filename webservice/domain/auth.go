package domain

type IAuthRepo interface {
	GetUserByUserId(UserId) (*User, error)
	GetUserByUsername(string) (*User, error)
	CreateUser(User) (*User, error)
	GrantUserClaim(UserId, Claim) (bool, error)
	StoreToken(UserId, TokenString) error
	RevokeToken(UserId) error
}
