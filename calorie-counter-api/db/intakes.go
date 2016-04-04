package db

import (
	"fmt"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/errors"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/models"
	"golang.org/x/net/context"
	"log"
	"time"
)

type IntakeRepository interface {
	Get(ctx context.Context, userID, intakeID string) (*models.Intake, error)
	GetAll(ctx context.Context, userID string) (models.Intakes, error)
	Create(ctx context.Context, intake *models.Intake) (*models.Intake, error)
	Update(ctx context.Context, intake *models.Intake) error
	Disable(ctx context.Context, intakeID, userID string) error
}

var Intakes IntakeRepository

type intakeRepository struct{}

func init() {
	Intakes = &intakeRepository{}
}

const intakeColumns = `id, user_id, name, calories, created_at, consumed_at, disabled_at`

func (i *intakeRepository) Get(ctx context.Context, userID, intakeID string) (*models.Intake, error) {
	log.Println("[intakes-get-one]", "UserID:", userID, "IntakeID:", intakeID)
	intake := models.Intake{}
	if _, err := db.QueryOne(&intake, fmt.Sprintf(`SELECT %s FROM intakes WHERE user_id=? AND id=? AND disabled_at IS NULL`, intakeColumns), userID, intakeID); err != nil {
		return nil, errors.New(errors.DATABASE_ERROR, "Error: %s", err.Error())
	}
	return &intake, nil
}

func (i *intakeRepository) GetAll(ctx context.Context, userID string) (models.Intakes, error) {
	log.Println("[intakes-get-all]", "UserID:", userID)
	intakes := make(models.Intakes, 0)
	if _, err := db.Query(&intakes, fmt.Sprintf(`SELECT %s FROM intakes WHERE user_id=? AND disabled_at IS NULL`, intakeColumns), userID); err != nil {
		return nil, errors.New(errors.DATABASE_ERROR, "Error: %s", err.Error())
	}
	return intakes, nil
}

func (i *intakeRepository) Create(ctx context.Context, intake *models.Intake) (*models.Intake, error) {
	log.Println("[intakes-create]", "UserID:", intake.UserID, "Name:", intake.Name, "Calories:", intake.Calories)
	if intake.ID == "" {
		intake.ID = NewUUID()
	} else if !IsUUID(intake.ID) {
		return nil, errors.New(errors.BAD_REQUEST, "Invalid intake ID: %s", intake.ID)
	}
	intake.CreatedAt = time.Now()
	intake.DisabledAt = nil

	_, err := db.ExecOne(`INSERT INTO intakes (id, user_id, name, calories, created_at, consumed_at, disabled_at)
	VALUES (?id, ?user_id, ?name, ?calories, now(), ?consumed_at, NULL)`, intake)
	if err != nil {
		return nil, errors.New(errors.DATABASE_ERROR, "Error: %s", err.Error())
	}
	return intake, nil
}

func (i *intakeRepository) Update(ctx context.Context, intake *models.Intake) error {
	log.Println("[intakes-update]", "ID:", intake.ID)

	query := `UPDATE intakes SET name=?name, calories=?calories, consumed_at=?consumed_at
		WHERE id = ?id AND disabled_at IS NULL`
	if _, err := db.ExecOne(query, intake); err != nil {
		return errors.New(errors.DATABASE_ERROR, "Error: %s", err.Error())
	}
	return nil
}

func (i *intakeRepository) Disable(ctx context.Context, intakeID, userID string) error {
	log.Println("[intakes-disable]", "ID:", intakeID, "UserID:", userID)
	if _, err := db.ExecOne(`UPDATE intakes SET disabled_at = now() WHERE id=? AND user_id=?`, intakeID, userID); err != nil {
		return errors.New(errors.DATABASE_ERROR, "Error: %s", err.Error())
	}
	return nil
}
