package controllers

import (
	"log"
	"net/http"

	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/models"
	"goji.io/pat"
	"golang.org/x/net/context"
)

func (a *Admin) GetUserIntake(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	log.Println("[ctrl-intakes-get]")
	req.ParseForm()
	userID := pat.Param(ctx, "user_id")
	intakeID := pat.Param(ctx, "intake_id")

	intake, err := db.Intakes.Get(ctx, userID, intakeID)
	if err != nil {
		ServeError(ctx, w, err)
		return
	}
	ServeJSON(ctx, w, intake)
}

func (a *Admin) GetUserIntakes(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	log.Println("[ctrl-intakes-get-for-user]")
	req.ParseForm()
	userID := pat.Param(ctx, "user_id")
	intakes, err := db.Intakes.GetAll(ctx, userID)
	if err != nil {
		ServeError(ctx, w, err)
		return
	}
	ServeJSON(ctx, w, models.IntakesData{Data: intakes})
}

func (a *Admin) CreateUserIntake(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	log.Println("[ctrl-intakes-create]")
	intake := &models.Intake{}
	ParseBody(ctx, intake, req)

	intake.UserID = pat.Param(ctx, "user_id")
	intake, err := db.Intakes.Create(ctx, intake)
	if err != nil {
		ServeError(ctx, w, err)
	}
	ServeJSON(ctx, w, intake)
}

func (a *Admin) UpdateUserIntake(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	intakeID := pat.Param(ctx, "intake_id")
	log.Println("[ctrl-intakes-update]", "ID:", intakeID)

	intake := &models.Intake{}
	ParseBody(ctx, intake, req)
	intake.ID = intakeID
	intake.UserID = pat.Param(ctx, "user_id")

	if err := db.Intakes.Update(ctx, intake); err != nil {
		ServeError(ctx, w, err)
	}
	ServeJSON(ctx, w, intake)
}

func (a *Admin) DisableUserIntake(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	intakeID := pat.Param(ctx, "intake_id")
	log.Println("[ctrl-intakes-disable]", "ID:", intakeID)

	userID := pat.Param(ctx, "user_id")
	if err := db.Intakes.Disable(ctx, intakeID, userID); err != nil {
		ServeError(ctx, w, err)
	}
	http.Redirect(w, req, "/", http.StatusOK)
}
