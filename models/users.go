package models

type User struct {
	Email string	`json:"email"`
	Password string	`json:"password"`
	Name 	string 	`json:"name"`
}

type Hello struct {
	Message string `json:"message"`
}