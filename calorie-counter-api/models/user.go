package models

type User struct {
	ID           string `json:"id"`
	Name         string `json:"name,omitempty"`
	Email        string `json:"email"`
	Password     string `json:"password,omitempty"`
	GoalCalories int    `json:"goalCalories"`
}

type Users []*User
