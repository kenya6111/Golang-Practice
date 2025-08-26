package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(username, email string) *User {
	return &User{
		ID:        uuid.NewString(),
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) ChangeUsername(newName string) {
	if newName != "" {
		u.Username = newName
		u.UpdatedAt = time.Now()
	}
}
