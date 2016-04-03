package controllers

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

type Admin struct{}

func (a *Admin) GetUsers(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Admin GetUsers")
}
