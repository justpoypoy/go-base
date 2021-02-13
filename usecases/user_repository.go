package usecases

import "github.com/justpoypoy/go-base/domain"

// UserRepository contract for user
type UserRepository interface {
	FindAll() (domain.Users, error)
	FindByID(userID uint) (domain.User, error)
}
