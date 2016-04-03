package controllers

import (
	"fmt"
	"goji.io/pat"
	"golang.org/x/net/context"
	"net/http"
)

type Intakes struct{}

func (a *Intakes) GetAll(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Intakes GetAll")
}

func (a *Intakes) Create(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Intakes Create")
}

func (a *Intakes) Update(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	intakeID := pat.Param(ctx, "id")
	fmt.Fprint(w, fmt.Sprintf("Intakes Update. ID: %s", intakeID))
}

func (a *Intakes) Disable(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	intakeID := pat.Param(ctx, "id")
	fmt.Fprint(w, fmt.Sprintf("Intakes Disable. ID: %s", intakeID))
}
