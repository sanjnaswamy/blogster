package types

// RegisterAuthorDTO defines author registration details
type RegisterAuthorDTO struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginDTO defines author login details
type LoginDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// UpdatePasswordDTO defines author password update details
type UpdatePasswordDTO struct {
	Password string `json:"password"`
}
