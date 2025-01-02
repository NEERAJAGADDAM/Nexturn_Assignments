package model

// Struct for user login requests
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Struct for the full user entity
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
