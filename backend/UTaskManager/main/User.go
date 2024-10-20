package main

type User struct {
	Id       int
	Username string
	Password string
	Name     string
	Surname  string
	Age      int
	Tasks    []Task // Tasks that user is doing now, has already completed or failed to do
	Roles    []Role // ROLE_USER, ROLE_ADMIN
}
