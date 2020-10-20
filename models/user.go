package models

type User interface {
	GetId() string
	GetName() string
}

type UserRepository interface {
	AddUser(user User)
	RemoveUser(user User)
	FindUserById(ID string) User
	GetAllUsers() []User
}
