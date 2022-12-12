package store

import "rest-pet/internal/model"

// UserRepository ....
type UserRepository interface {
	Create(user *model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
