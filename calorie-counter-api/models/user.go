package models

import "time"

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

type User struct {
	ID           string     `json:"id"`
	Email        string     `json:"email"`
	Password     string     `json:"password,omitempty"`
	Role         string     `json:"role"`
	Name         string     `json:"name,omitempty"`
	GoalCalories int        `json:"goalCalories"`
	DisableAt    *time.Time `json:"disabledAt,omitempty"`
}

type Users []*User

// NewRecord implements factory method needed for DB queries
func (users *Users) NewRecord() interface{} {
	user := &User{}
	*users = append(*users, user)
	return user
}

type UsersData struct {
	Data Users `json:"users"`
}
