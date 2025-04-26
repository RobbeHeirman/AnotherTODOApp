package persistence

type Repository interface {
	CreateUser(user *User) error
}
