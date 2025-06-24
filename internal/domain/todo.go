package domain

import (
	"time"
)

type Todo struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	UserID uint `json:"user_id" gorm:"not null;index"`
	User   User `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type TodoRepository interface {
	Create(todo *Todo) error
	GetByID(id uint) (*Todo, error)
	GetByUserID(userID uint) ([]Todo, error)
	Update(todo *Todo) error
	Delete(id uint) error
}

type TodoService interface {
	CreateTodo(userID uint, title, description string) (*Todo, error)
	GetTodo(id uint) (*Todo, error)
	GetTodosByUser(userID uint) ([]Todo, error)
	UpdateTodo(id uint, title, description string, completed bool) (*Todo, error)
	DeleteTodo(id uint) error
	ToggleTodoStatus(id uint) (*Todo, error)
}
