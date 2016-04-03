package db

import "github.com/FcoManueel/calorie-counter/calorie-counter-api/models"

type IntakeRepository interface {
	GetAll(userID string) (models.Intakes, error)
}

var Intakes IntakeRepository

type intakeRepository struct{}

func init() {
	Intakes = &intakeRepository{}
}

func (i *intakeRepository) GetAll(userID string) (models.Intakes, error) {
	return nil, nil
}
