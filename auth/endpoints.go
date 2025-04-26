package auth

import "net/http"

type User struct {
	Name     string `json:"user"`
	Password string `json:"password"`
}

func Register(user *User, w http.ResponseWriter, r *http.Request) any {
	return struct {
		response string
	}{
		"Done",
	}
}
