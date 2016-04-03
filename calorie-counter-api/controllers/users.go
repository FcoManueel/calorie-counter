package controllers

import (
	"golang.org/x/net/context"
	"net/http"
	"time"

	"errors"
	"fmt"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/models"
	"goji.io/pat"
)

type Users struct{}

func (a *Users) Get(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	//	user := &models.User
	//	ParseBody(user, req)
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
	ParseBody(user, req)
	userID := pat.Param(ctx, "id")

	if err := db.Users.Update("UPDATE users SET name='Marciano Pajarito'", userID); err != nil {
		ServeError(ctx, w, errors.New(fmt.Sprintf("Error while updating user. Error: %s", err.Error())))
	}
	ServeJSON(ctx, w, user)
}

func (a *Users) Disable(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	ParseBody(user, req)
	now := time.Now()
	user.DisableAt = &now
	http.Redirect(w, req, "/", http.StatusOK)
}
