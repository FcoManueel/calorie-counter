package models

import "time"

type Intake struct {
	ID         string    `json:"id"`
	UserID     string    `json:"userId"`
	Name       string    `json:"name"`
	Calories   int       `json:"calories"`
	CreatedAt  time.Time `json:"createdAt"`
	ConsumedAt time.Time `json:"consumedAt"`
}

type Intakes []*Intake
