package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string
	Username  string
	Email     Email
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(username, email string) (*User, error) {
	mail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        uuid.NewString(),
		Username:  username,
		Email:     mail,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (u *User) ChangeUsername(newName string) {
	if newName != "" {
		u.Username = newName
		u.UpdatedAt = time.Now()
	}
}
