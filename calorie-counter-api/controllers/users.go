package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/models"
	"goji.io/pat"
	"golang.org/x/net/context"
)

type Users struct{}

func (a *Users) Get(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	userID := pat.Param(ctx, "id")
	user, err := db.Users.Get(userID)
	if err != nil {
		ServeError(ctx, w, err)
		return
	}
	ServeJSON(ctx, w, user)
}

func (a *Users) Update(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	userID := pat.Param(ctx, "id")
	ParseBody(user, req)
	user.ID = userID

	if err := db.Users.Update(user); err != nil {
		ServeError(ctx, w, errors.New(fmt.Sprintf("Error while updating user. ID: %s, Error: %s", user.ID, err.Error())))
	}
	ServeJSON(ctx, w, user)
}

func (a *Users) Disable(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	userID := pat.Param(ctx, "id")
	if err := db.Users.Disable(userID); err != nil {
		ServeError(ctx, w, errors.New(fmt.Sprintf("Error while disabling user. ID: %s,  Error: %s", userID, err.Error())))
	}
	http.Redirect(w, req, "/", http.StatusOK)
}
