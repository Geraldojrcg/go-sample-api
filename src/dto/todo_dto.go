package dto

type CreateTodoDto struct {
	Description string `validate:"required" json:"description"`
	UserID      string `validate:"required,uuid4" json:"user_id"`
}

type UpdateTodoDto struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
