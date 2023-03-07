package model

import (
	ttv1 "github.com/borud/tt/pkg/tt/v1"
)

// User represents a user.
type User struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
}

// Proto returns proto representation.
func (u User) Proto() *ttv1.User {
	return &ttv1.User{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Phone:    u.Phone,
	}
}

// UserFromProto returns User corresponding to ttv1.User
func UserFromProto(u *ttv1.User) User {
	return User{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Phone:    u.Phone,
	}
}
