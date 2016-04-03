package db

import (
	"errors"
	"fmt"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/models"
	"log"
	"strings"
)

type UserRepository interface {
	Get(userID string) (*models.User, error)
	GetAll() (models.Users, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) error
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
		return nil, errors.New(fmt.Sprintf("No user found. ID: %s", userID))
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
	} else if !IsUUID(user.ID) {
		return nil, errors.New(fmt.Sprintf("Invalid user ID: %s", user.ID))
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

func (u *userRepository) Update(user *models.User) (err error) {
	var updateFields []string
	if user.Name != "" {
		updateFields = append(updateFields, "name=?name")
	}
	if user.Email != "" {
		updateFields = append(updateFields, "email=?email")
	}
	if user.Password != "" {
		if user.Password, err = Hash(user.Password, []byte(user.ID)); err != nil {
			return errors.New(fmt.Sprintf("Error on user update. ID: %s Error: %s", user.ID, err.Error()))
		}
		updateFields = append(updateFields, "password=?password")
	}
	if user.GoalCalories != 0 {
		updateFields = append(updateFields, "goal_calories=?goal_calories")
	}

	if len(updateFields) > 0 {
		query := fmt.Sprintf(`UPDATE users SET %s WHERE id = ?id AND disabled_at IS NULL`, strings.Join(updateFields, ","))
		if _, err := db.ExecOne(query, user); err != nil {
			return errors.New(fmt.Sprintf("Error on user update. ID: %s Error: %s", user.ID, err.Error()))
		}
	}
	return nil
}

func (u *userRepository) Disable(userID string) error {
	if _, err := db.ExecOne(`UPDATE users SET disabled_at = now() WHERE id = ?`, userID); err != nil {
		return errors.New(fmt.Sprintf("Error while disabling user. ID: %s. Error: %s", userID, err.Error()))
	}
	return nil
}
