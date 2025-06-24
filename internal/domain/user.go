package domain

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex;not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Todos []Todo `json:"todos,omitempty" gorm:"foreignKey:UserID"`
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() ([]User, error)
	Update(user *User) error
	Delete(id uint) error
}

type UserService interface {
	CreateUser(username, email, password, name string) (*User, error)
	GetUser(id uint) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(id uint, username, email, name string) (*User, error)
	DeleteUser(id uint) error
	GetUserWithTodos(id uint) (*User, error)
}
