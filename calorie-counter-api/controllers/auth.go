package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/FcoManueel/calorie-counter/calorie-counter-api/db"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/errors"
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
	ParseBody(ctx, user, req)

	user.Role = models.RoleUser
	var err error
	if user, err = db.Users.Create(ctx, user); err != nil {
		ServeError(ctx, w, err)
		return
	}

	ServeJSON(ctx, w, user)
}

func (a *Auth) Login(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	form := &models.SignInForm{}
	ParseBody(ctx, form, req)
	authToken, err := IssueToken(ctx, form.Email, form.Password)
	if err != nil {
		ServeError(ctx, w, err)
		return
	}
	ServeJSON(ctx, w, authToken)
}

func IssueToken(ctx context.Context, email, password string) (*models.AuthToken, error) {
	user, err := db.Users.GetByEmail(ctx, strings.ToLower(email))
	if err != nil {
		if err, ok := err.(errors.AppError); ok && err.Code != errors.INTERNAL_SERVER_ERROR {
			return nil, errors.New(errors.UNAUTHORIZED, "Wrong credentials")
		}
		return nil, err
	}
	if user.DisableAt != nil {
		return nil, errors.New(errors.FORBIDDEN, "User has been disabled")
	}
	var hashedPassword string
	if hashedPassword, err = db.Hash(password, []byte(user.ID)); err != nil {
		return nil, errors.New(errors.INTERNAL_SERVER_ERROR, "Error while hashing: %s", err.Error())
	}
	if user.Password != hashedPassword {
		return nil, errors.New(errors.UNAUTHORIZED, "Wrong credentials")
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
