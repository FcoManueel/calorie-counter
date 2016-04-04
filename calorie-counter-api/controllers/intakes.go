package controllers

import (
	"log"
	"net/http"

	"github.com/FcoManueel/calorie-counter/calorie-counter-api/ccontext"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/models"
	"goji.io/pat"
	"golang.org/x/net/context"
)

type Intakes struct{}

func (in *Intakes) Get(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	log.Println("[ctrl-intakes-get]")
	req.ParseForm()
	userID := ccontext.GetUserID(ctx)
	intakeID := pat.Param(ctx, "id")

	intake, err := db.Intakes.Get(ctx, userID, intakeID)
	if err != nil {
		ServeError(ctx, w, err)
		return
	}
	ServeJSON(ctx, w, intake)
}

func (in *Intakes) GetAll(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	log.Println("[ctrl-intakes-get-for-user]")
	req.ParseForm()
	userID := ccontext.GetUserID(ctx)
	intakes, err := db.Intakes.GetAll(ctx, userID)
	if err != nil {
		ServeError(ctx, w, err)
		return
	}
	ServeJSON(ctx, w, models.IntakesData{Data: intakes})
}

func (in *Intakes) Create(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	log.Println("[ctrl-intakes-create]")
	intake := &models.Intake{}
	ParseBody(ctx, intake, req)

	intake.UserID = ccontext.GetUserID(ctx)
	intake, err := db.Intakes.Create(ctx, intake)
	if err != nil {
		ServeError(ctx, w, err)
	}
	ServeJSON(ctx, w, intake)
}

func (in *Intakes) Update(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	intakeID := pat.Param(ctx, "id")
	log.Println("[ctrl-intakes-update]", "ID:", intakeID)

	intake := &models.Intake{}
	ParseBody(ctx, intake, req)
	intake.ID = intakeID
	intake.UserID = ccontext.GetUserID(ctx)

	if err := db.Intakes.Update(ctx, intake); err != nil {
		ServeError(ctx, w, err)
	}
	ServeJSON(ctx, w, intake)
}

func (in *Intakes) Disable(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	intakeID := pat.Param(ctx, "id")
	log.Println("[ctrl-intakes-disable]", "ID:", intakeID)

	userID := ccontext.GetUserID(ctx)
	if err := db.Intakes.Disable(ctx, intakeID, userID); err != nil {
		ServeError(ctx, w, err)
	}
	http.Redirect(w, req, "/", http.StatusOK)
}
