package controllers

import (
	"fmt"
	"goji.io/pat"
	"golang.org/x/net/context"
	"net/http"
)

type Admin struct{}

func (a *Admin) GetUsers(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Admin GetUsers")
}

func (a *Admin) CreateUser(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Admin CreateUser")
}
func (a *Admin) GetUser(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	userID := pat.Param(ctx, "user_id")
	fmt.Fprint(w, "Admin GetUser. ID: %s", userID)
}
func (a *Admin) UpdateUser(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	userID := pat.Param(ctx, "user_id")
	fmt.Fprint(w, "Admin UpdateUser. ID: %s", userID)
}
func (a *Admin) DisableUser(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	userID := pat.Param(ctx, "user_id")
	fmt.Fprint(w, "Admin DisableUser. ID: %s", userID)
}
