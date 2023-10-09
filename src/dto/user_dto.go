package dto

type CreateUserDto struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UpdateUserDto struct {
	Name  string `json:"name"`
	Email string `validate:"omitempty,email" json:"email"`
}
