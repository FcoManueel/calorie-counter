package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/models"
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"golang.org/x/net/context"
)

type Auth struct{}

const tokenExpiration = 10 * time.Hour

func (a *Auth) Signup(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	ParseBody(user, req)

	user.Role = db.RoleUser
	var err error
	if user, err = db.Users.Create(user); err != nil {
		ServeError(ctx, w, errors.New(fmt.Sprintf("Error on signup. Error: %s", err.Error())))
		return
	}

	ServeJSON(ctx, w, user)
}

func (a *Auth) Login(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	form := &models.SignInForm{}
	ParseBody(form, req)
	authToken, err := IssueToken(ctx, form.Email, form.Password)
	if err != nil {
		ServeError(ctx, w, errors.New(fmt.Sprintf("Login error: %s", err.Error())))
		return
	}
	ServeJSON(ctx, w, authToken)
}

func IssueToken(ctx context.Context, email, password string) (*models.AuthToken, error) {
	user, err := db.Users.GetByEmail(strings.ToLower(email))
	if err != nil {
		return nil, err
	}
	if user.DisableAt != nil {
		return nil, errors.New("Disabled user")
	}
	var hashedPassword string
	if hashedPassword, err = db.Hash(password, []byte(user.ID)); err != nil {
		return nil, errors.New(fmt.Sprintf("Error while hashing: %s", err.Error()))
	}
	if user.Password != hashedPassword {
		return nil, errors.New("Wrong password")
	}

	tokenID := db.NewUUID()
	jwt, err := createAuthToken(ctx, tokenID, user.ID, user.Role)
	if err != nil {
		return nil, err
	}
	authToken, err := jwt.Serialize([]byte(db.JWTSigningKey))
	if err != nil {
		return nil, err
	}
	return &models.AuthToken{AccessToken: string(authToken)}, nil
}

// CreateAuthToken generates a JWT access token
func createAuthToken(ctx context.Context, tokenID, userID, scope string) (jwt.JWT, error) {
	now := time.Now()
	expiry := now.Add(tokenExpiration).Unix()

	claims := jws.Claims{}
	claims.SetSubject(userID)
	claims.SetJWTID(tokenID)
	claims.SetIssuedAt(float64(now.Unix()))
	claims.SetExpiration(float64(expiry))
	claims.Set("scope", scope)
	jwt := jws.NewJWT(claims, crypto.SigningMethodHS256)
	return jwt, nil
}
