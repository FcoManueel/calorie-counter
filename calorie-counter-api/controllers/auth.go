package controllers

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

type Auth struct{}

func (a *Auth) Login(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	http.Error(w, "WOLOLO_ERROR", 500)
	fmt.Fprint(w, "Auth Login")
}
