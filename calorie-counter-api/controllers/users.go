package controllers

import (
	"net/http"

	"github.com/FcoManueel/calorie-counter/calorie-counter-api/ccontext"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/models"
	"goji.io/pat"
	"golang.org/x/net/context"
)

type Users struct{}

func (a *Users) Get(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	user, err := db.Users.Get(ctx, ccontext.GetUserID(ctx))
	if err != nil {
		ServeError(ctx, w, err)
		return
	}
	ServeJSON(ctx, w, user)
}

func (a *Users) Update(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	userID := pat.Param(ctx, "id")
	ParseBody(ctx, user, req)
	user.ID = userID

	if err := db.Users.Update(ctx, user); err != nil {
		ServeError(ctx, w, err)
		return
	}
	ServeJSON(ctx, w, user)
}

func (a *Users) Disable(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	userID := pat.Param(ctx, "id")
	if err := db.Users.Disable(ctx, userID); err != nil {
		ServeError(ctx, w, err)
		return
	}
	http.Redirect(w, req, "/", http.StatusOK)
}
