package middleware

import (
	"errors"
	"fmt"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/ccontext"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/controllers"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/models"
	"goji.io"
	"golang.org/x/net/context"
	"net/http"
)

func RequireAdmin(h goji.Handler) goji.Handler {
	handler := func(ctx context.Context, w http.ResponseWriter, req *http.Request) {
		if ccontext.GetRole(ctx) != models.RoleAdmin {
			controllers.ServeError(ctx, w, errors.New(fmt.Sprintf("Not an admin. Role: %s", ccontext.GetRole(ctx))))
			return
		}
		h.ServeHTTPC(ctx, w, req)
	}
	return goji.HandlerFunc(handler)
}
