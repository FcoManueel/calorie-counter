package controllers

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

type Users struct{}

func (a *Users) Get(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Users Get")
}
