package controllers

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

type Intakes struct{}

func (a *Intakes) GetAll(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Intakes GetAll")
}
