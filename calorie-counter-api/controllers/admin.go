package controllers

import (
	"errors"
	"fmt"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/models"
	"goji.io/pat"
	"golang.org/x/net/context"
	"net/http"
)

type Admin struct{}

func (a *Admin) GetUsers(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	users, err := db.Users.GetAll(ctx)
	if err != nil {
		ServeError(ctx, w, err)
		return
	}
	ServeJSON(ctx, w, models.UsersData{Data: users})
}

func (a *Admin) CreateUser(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	ParseBody(ctx, user, req)

	if user.Role == "" {
		user.Role = models.RoleUser
	}
	var err error
	if user, err = db.Users.Create(ctx, user); err != nil {
		ServeError(ctx, w, errors.New(fmt.Sprintf("Error on signup. Error: %s", err.Error())))
		return
	}

	ServeJSON(ctx, w, user)
}
func (a *Admin) GetUser(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	userID := pat.Param(ctx, "user_id")
	user, err := db.Users.Get(ctx, userID)
	if err != nil {
		ServeError(ctx, w, err)
		return
	}
	ServeJSON(ctx, w, user)
}

func (a *Admin) UpdateUser(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	userID := pat.Param(ctx, "user_id")
	ParseBody(ctx, user, req)
	user.ID = userID

	if err := db.Users.Update(ctx, user); err != nil {
		ServeError(ctx, w, errors.New(fmt.Sprintf("Error while updating user. ID: %s, Error: %s", user.ID, err.Error())))
		return
	}
	ServeJSON(ctx, w, user)
}

func (a *Admin) DisableUser(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	userID := pat.Param(ctx, "user_id")
	if err := db.Users.Disable(ctx, userID); err != nil {
		ServeError(ctx, w, errors.New(fmt.Sprintf("Error while disabling user. ID: %s,  Error: %s", userID, err.Error())))
		return
	}
	http.Redirect(w, req, "/", http.StatusOK)
}
