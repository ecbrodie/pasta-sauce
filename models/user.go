package models

type User struct {
	Name string
}

func NewUser(name string) *User {
	return &User{name}
}
