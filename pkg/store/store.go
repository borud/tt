package store

import "github.com/borud/tt/pkg/model"

// Store interface for tt
type Store interface {
	Close() error

	AddUser(user model.User) error
	GetUser(username string) (model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(username string) error
	ListUsers() ([]model.User, error)
}
