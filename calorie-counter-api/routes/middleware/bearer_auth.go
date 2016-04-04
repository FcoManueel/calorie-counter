package middleware

import (
	"log"
	"net/http"

	"github.com/FcoManueel/calorie-counter/calorie-counter-api/ccontext"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/controllers"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/errors"
	"github.com/dgrijalva/jwt-go"
	"goji.io"
	"golang.org/x/net/context"
)

func BearerAuth(h goji.Handler) goji.Handler {
	handler := func(ctx context.Context, w http.ResponseWriter, req *http.Request) {
		ctx, err := validateBearerToken(ctx, req)
		if err != nil {
			controllers.ServeError(ctx, w, err)
			return
		}
		h.ServeHTTPC(ctx, w, req)
	}
	return goji.HandlerFunc(handler)
}

func validateBearerToken(ctx context.Context, req *http.Request) (context.Context, error) {
	token, err := jwt.ParseFromRequest(req, func(*jwt.Token) (interface{}, error) { return []byte(db.JWTSigningKey), nil })
	if err != nil {
		return nil, err
	}
	userID := token.Claims["sub"].(string)
	if !db.IsUUID(userID) {
		return nil, errors.New(errors.BAD_REQUEST, "Invalid userID in bearerToken")
	}
	user, err := db.Users.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user.DisableAt != nil {
		return nil, errors.New(errors.FORBIDDEN, "Disabled user")
	}
	role := token.Claims["scope"].(string)
	ctx = ccontext.SetRole(ctx, role)
	ctx = ccontext.SetUserID(ctx, userID)
	log.Println(ctx, " (bearer-auth) Ok", "userID", userID, "role", role)
	return ctx, nil
}
