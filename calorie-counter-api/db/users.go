package db

import (
	"fmt"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/errors"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/models"
	"golang.org/x/net/context"
	"log"
	"strings"
)

type UserRepository interface {
	Get(ctx context.Context, userID string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetAll(ctx context.Context) (models.Users, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Disable(ctx context.Context, userID string) error
}

var Users UserRepository

type userRepository struct{}

func init() {
	Users = &userRepository{}
}

const userColumns = `id, role, name, email, password, goal_calories`

func (u *userRepository) Get(ctx context.Context, userID string) (*models.User, error) {
	log.Println("[user-get]", "ID:", userID)
	if !IsUUID(userID) {
		return nil, errors.New(errors.BAD_REQUEST, "Invalid user ID: %s", userID)
	}

	user := &models.User{}
	if _, err := db.QueryOne(user, fmt.Sprintf(`SELECT %s FROM users WHERE id=?`, userColumns), userID); err != nil {
		return nil, errors.New(errors.DATABASE_ERROR, "Error: %s", err.Error())
	}
	user.Password = ""
	return user, nil
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	log.Println("[user-get]", "Email:", email)

	user := &models.User{}
	if _, err := db.QueryOne(user, fmt.Sprintf(`SELECT %s FROM users WHERE email = ?`, userColumns), email); err != nil {
		return nil, errors.New(errors.DATABASE_ERROR, "Error: %s", err.Error())
	}
	return user, nil
}

func (u *userRepository) GetAll(ctx context.Context) (models.Users, error) {
	log.Println("[user-get-all]")
	users := make(models.Users, 0)
	if _, err := db.Query(&users, fmt.Sprintf(`SELECT %s FROM users WHERE disabled_at IS NULL`, userColumns)); err != nil {
		return nil, errors.New(errors.DATABASE_ERROR, "Error: %s", err.Error())
	}
	for i := range users {
		users[i].Password = ""
	}
	return users, nil
}

func (u *userRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	log.Println("[user-create]", "Email:", user.Email)
	if user.ID == "" {
		user.ID = NewUUID()
	} else if !IsUUID(user.ID) {
		return nil, errors.New(errors.BAD_REQUEST, "Invalid user ID: %s", user.ID)
	}
	log.Println("No password")
	if user.Password != "" {
		log.Println("Password", user.Password)
		key, err := Hash(user.Password, []byte(user.ID))
		if err != nil {
			return nil, errors.New(errors.INTERNAL_SERVER_ERROR, "Error while hashing: %s", err.Error())
		}
		user.Password = key
	}
	user.Email = strings.ToLower(user.Email)

	_, err := db.ExecOne(`INSERT INTO users (id, role, name, email, password, goal_calories)
	VALUES (?id, ?role, ?name, ?email, ?password, ?goal_calories)`, user)
	if err != nil {
		return nil, errors.New(errors.DATABASE_ERROR, "Error: %s", err.Error())
	}
	return user, nil
}

func (u *userRepository) Update(ctx context.Context, user *models.User) (err error) {
	log.Println("[user-update]", "ID:", user.ID)
	var updateFields []string
	if user.Name != "" {
		updateFields = append(updateFields, "name=?name")
	}
	if user.Email != "" {
		user.Email = strings.ToLower(user.Email)
		updateFields = append(updateFields, "email=?email")
	}
	if user.Password != "" {
		if user.Password, err = Hash(user.Password, []byte(user.ID)); err != nil {
			return errors.New(errors.INTERNAL_SERVER_ERROR, "Error while hashing: %s", err.Error())
		}
		updateFields = append(updateFields, "password=?password")
	}
	if user.GoalCalories != 0 {
		updateFields = append(updateFields, "goal_calories=?goal_calories")
	}

	if user.Role != "" {
		updateFields = append(updateFields, "role=?role")
	}

	if len(updateFields) > 0 {
		query := fmt.Sprintf(`UPDATE users SET %s WHERE id = ?id AND disabled_at IS NULL`, strings.Join(updateFields, ","))
		if _, err := db.ExecOne(query, user); err != nil {
			return errors.New(errors.DATABASE_ERROR, "Error: %s", err.Error())
		}
	}
	return nil
}

func (u *userRepository) Disable(ctx context.Context, userID string) error {
	log.Println("[user-disable]", "ID:", userID)
	if _, err := db.ExecOne(`UPDATE users SET disabled_at = now() WHERE id = ?`, userID); err != nil {
		return errors.New(errors.DATABASE_ERROR, "Error: %s", err.Error())
	}
	return nil
}
