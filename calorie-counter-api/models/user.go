package models

import "time"

type User struct {
	ID           string     `json:"id"`
	Email        string     `json:"email"`
	Password     string     `json:"password,omitempty"`
	Role         string     `json:"role,omitempty"`
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
