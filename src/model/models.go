package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:'uuid_generate_v4()'" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type User struct {
	BaseModel `gorm:"embedded"`

	Name     string `json:"name"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"-"`

	Todos []Todo `json:"todos,omitempty"`
}

type Todo struct {
	BaseModel `gorm:"embedded"`

	Description string `json:"description"`
	Completed   bool   `json:"completed"`

	UserID uuid.UUID `json:"-"`
	User   User      `json:"user"`
}
