package models

type User struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type UserId struct {
	Id int
}

type UserLoggedIn struct {
	AccessToken string `json:"accessToken"`
}

type UserLogsInDb struct {
	Id       int    `db:"id"`
	Password string `db:"password"`
}
