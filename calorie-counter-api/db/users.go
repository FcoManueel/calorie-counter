package db

import (
	"errors"
	"fmt"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/models"
	"log"
)

type UserRepository interface {
	Get(userID string) (*models.User, error)
	GetAll() (models.Users, error)
	Create(user *models.User) (*models.User, error)
	Update(query, userID string) error
	Disable(userID string) error
}

var Users UserRepository

type userRepository struct{}

func init() {
	Users = &userRepository{}
}

const userColumns = `id, role, name, email, password, goal_calories`
const RoleUser = "user"

func (u *userRepository) Get(userID string) (*models.User, error) {
	if !IsUUID(userID) {
		return nil, errors.New(fmt.Sprintf("Invalid user ID: %s", userID))
	}

	user := &models.User{ID: userID}
	if _, err := db.QueryOne(user, fmt.Sprintf(`SELECT %s FROM users WHERE id = ?id`, userColumns), user); err != nil {
		return nil, errors.New(fmt.Sprintf("No user found having id: %s", userID))
	}
	return user, nil
}

func (u *userRepository) GetAll() (models.Users, error) {
	users := make(models.Users, 0)
	if _, err := db.Query(&users, fmt.Sprintf(`SELECT %s FROM users WHERE disabled != NULL`, userColumns)); err != nil {
		return nil, errors.New(fmt.Sprintf("Error while fetching users: %s", err.Error()))
	}
	return users, nil
}

func (u *userRepository) Create(user *models.User) (*models.User, error) {
	log.Println("[user-create]", "Email:", user.Email)
	if user.ID == "" {
		user.ID = NewUUID()
	}
	if user.Password != "" {
		key, err := Hash(user.Password, []byte(user.ID))
		if err != nil {
			return nil, err
		}
		user.Password = key
	}
	_, err := db.ExecOne(`INSERT INTO users (id, role, name, email, password, goal_calories)
	VALUES (?id, ?role, ?name, ?email, ?password, ?goal_calories)`, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) Update(query, userID string) error {
	query = fmt.Sprintf(`%s WHERE id = ?`, query)
	if _, err := db.ExecOne(query, userID); err != nil {
		return errors.New(fmt.Sprintf("Error on user update. ID: %s Error: %s", userID, err.Error()))
	}
	return nil
}

func (u *userRepository) Disable(userID string) error {
	user := &models.User{ID: userID}
	if _, err := db.ExecOne(`UPDATE users SET disabled_at = now() WHERE id = ?id`, user); err != nil {
		return errors.New(fmt.Sprintf("Error while disabling user with id: %s. Error: %s", userID, err.Error()))
	}
	return nil
}
