package models

// User Model
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Page Model
type Page struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Link   string `json:"link"`
	UserID string `json:"user_id"`
}

// Project Model
type Project struct {
	ID       string `json:"id"`
	Name     string `json:"description"`
	Link     string `json:"link"`
	TestLink string `json:"testlink"`
	UserID   string `json:"user_id"`
}

// Guide Model
type Guide struct {
	ID     string `json:"id"`
	Name   string `json:"description"`
	Link   string `json:"link"`
	UserID string `json:"user_id"`
}
