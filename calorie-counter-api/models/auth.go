package models

type AuthToken struct {
	AccessToken string `json:"accessToken"`
}

type SignInForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
