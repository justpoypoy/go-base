package usecases

import "github.com/justpoypoy/go-base/domain"

// A UserInteractor belong to the usecases layer.
type UserInteractor struct {
	UserRepository UserRepository
}

// Index is display a listing of the resource.
func (ui *UserInteractor) Index() (domain.Users, error) {
	users, err := ui.UserRepository.FindAll()
	return users, err
}

// Show is display the specified resource.
func (ui *UserInteractor) Show(userID uint) (domain.User, error) {
	user, err := ui.UserRepository.FindByID(userID)
	return user, err
}
