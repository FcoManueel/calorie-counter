package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/FcoManueel/calorie-counter/calorie-counter-api/controllers"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/dgrijalva/jwt-go"
	"goji.io"
	"golang.org/x/net/context"
)

func BearerAuth(h goji.Handler) goji.Handler {
	handler := func(ctx context.Context, w http.ResponseWriter, req *http.Request) {
		if err := validateBearerToken(ctx, req); err != nil {
			controllers.ServeError(ctx, w, err)
			return
		}
		h.ServeHTTPC(ctx, w, req)
	}
	return goji.HandlerFunc(handler)
}

func validateBearerToken(ctx context.Context, req *http.Request) error {
	token, err := jwt.ParseFromRequest(req, func(*jwt.Token) (interface{}, error) { return []byte(db.JWTSigningKey), nil })
	if err != nil {
		return err
	}
	userID := token.Claims["sub"].(string)
	if !db.IsUUID(userID) {
		return errors.New("Invalid userID in bearerToken")
	}
	user, err := db.Users.Get(ctx, userID)
	if err != nil {
		return err
	}
	if user.DisableAt != nil {
		return errors.New("Disabled user")
	}
	ctx = context.WithValue(ctx, "userID", userID)
	role := token.Claims["scope"].(string)
	ctx = context.WithValue(ctx, "role", role)
	log.Println(ctx, " (bearer-auth) Ok", "userID", userID, "role", role)
	return nil
}
