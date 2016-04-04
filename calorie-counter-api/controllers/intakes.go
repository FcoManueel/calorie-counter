package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/models"
	"goji.io/pat"
	"golang.org/x/net/context"
)

type Intakes struct{}

const stubUserID = "3ee205c7-9333-4edd-9e90-429f3a040259"

func (a *Intakes) Get(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	log.Println("[ctrl-intakes-get]")
	req.ParseForm()
	userID := stubUserID
	intakeID := pat.Param(ctx, "id")

	intake, err := db.Intakes.Get(ctx, userID, intakeID)
	if err != nil {
		ServeError(ctx, w, err)
		return
	}
	ServeJSON(ctx, w, intake)
}

func (a *Intakes) GetAll(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	log.Println("[ctrl-intakes-get-for-user]")
	req.ParseForm()
	userID := stubUserID
	intakes, err := db.Intakes.GetAll(ctx, userID)
	if err != nil {
		ServeError(ctx, w, err)
		return
	}
	ServeJSON(ctx, w, models.IntakesData{Data: intakes})
}

func (a *Intakes) Create(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	log.Println("[ctrl-intakes-create]")
	intake := &models.Intake{}
	ParseBody(ctx, intake, req)

	intake.UserID = stubUserID
	intake, err := db.Intakes.Create(ctx, intake)
	if err != nil {
		ServeError(ctx, w, errors.New(fmt.Sprintf("Error while creating intake. Error: %s", err.Error())))
	}
	ServeJSON(ctx, w, intake)
}

func (a *Intakes) Update(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	intakeID := pat.Param(ctx, "id")
	log.Println("[ctrl-intakes-update]", "ID:", intakeID)

	intake := &models.Intake{}
	ParseBody(ctx, intake, req)
	intake.ID = intakeID
	intake.UserID = stubUserID

	if err := db.Intakes.Update(ctx, intake); err != nil {
		ServeError(ctx, w, errors.New(fmt.Sprintf("Error while updating intake. ID: %s, Error: %s", intake.ID, err.Error())))
	}
	ServeJSON(ctx, w, intake)
}

func (a *Intakes) Disable(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	intakeID := pat.Param(ctx, "id")
	log.Println("[ctrl-intakes-disable]", "ID:", intakeID)

	userID := stubUserID
	if err := db.Intakes.Disable(ctx, intakeID, userID); err != nil {
		ServeError(ctx, w, errors.New(fmt.Sprintf("Error while disabling intake. ID: %s, Error: %s", intakeID, err.Error())))
	}
	http.Redirect(w, req, "/", http.StatusOK)

}
