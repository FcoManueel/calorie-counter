package main

import (
	"fmt"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

func init() {
	if err := db.Migrate("gouser", "gouser", "127.0.0.1", "calories_dev"); err != nil {
		log.Fatal("[db-migrate] Error while applying migrations: ", err.Error())
	}
}

func hello(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(ctx, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/hello/:name"), hello)

	http.ListenAndServe("localhost:8000", mux)
}
