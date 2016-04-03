package db

import "github.com/FcoManueel/calorie-counter/calorie-counter-api/models"

type UserRepository interface {
	Get(userID string) (*models.User, error)
	GetAll() (models.Users, error)
	Create(user *models.User) (*models.User, error)
	Update(userID string) (*models.User, error)
	Disable(userID string) error
}

var Users UserRepository

type userRepository struct{}

func init() {
	Users = &userRepository{}
}

func (u *userRepository) Get(userID string) (*models.User, error) {
	return nil, nil
}

func (u *userRepository) GetAll() (models.Users, error) {
	return nil, nil
}

func (u *userRepository) Create(user *models.User) (*models.User, error) {
	return nil, nil
}

func (u *userRepository) Update(userID string) (*models.User, error) {
	return nil, nil
}

func (u *userRepository) Disable(userID string) error {
	return nil
}
