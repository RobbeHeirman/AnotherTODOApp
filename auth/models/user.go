package models

type User struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type UserId struct {
	Id int
}
