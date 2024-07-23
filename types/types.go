package types

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserStore interface {
	CreateUser(User) error
	GetUserByID(int) (*User, error)
	GetUserByUsernameOrEmail(string) (*User, error)
}
