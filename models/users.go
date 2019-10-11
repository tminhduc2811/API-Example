package models

type User struct {
	Email string	`json:"email"`
	Password string	`json:"password,omitempty"`
	Name 	string 	`json:"name"`
	HashPassword []byte        `json:"hashpassword,omitempty"`
}

type Hello struct {
	Message string `json:"message"`
}