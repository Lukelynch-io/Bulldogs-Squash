package domain

type IAuthRepo interface {
	GetUserByUserId(UserId) (*User, error)
	GetUserByUsername(string) (*User, error)
	CreateUser(User) (bool, error)
	GrantUserClaim(UserId, Claim) (bool, error)
}
