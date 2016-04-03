package controllers

import (
	"errors"
	"fmt"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/models"
	"golang.org/x/net/context"
	"net/http"
)

type Auth struct{}

func (a *Auth) Signup(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	ParseBody(user, req)

	user.Role = db.RoleUser
	var err error
	if user, err = db.Users.Create(user); err != nil {
		ServeError(ctx, w, errors.New(fmt.Sprintf("Error on signup. Error: %s", err.Error())))
		return
	}

	ServeJSON(ctx, w, user)
}

func (a *Auth) Login(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Auth Login")
}
