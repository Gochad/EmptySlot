package internal

import (
	"os"
)

type config struct {
	TokenSecret    string
	RedirectURL    string
	ClientID       string
	AuthSecret     string
	Oauth2URL      string
	DbUser         string
	DbPwd          string
	DbName         string
	TokenSecretJWT string
}

var EnvConfig config

func init() {
	EnvConfig = config{
		TokenSecret:    os.Getenv("TOKEN_SECRET"),
		RedirectURL:    os.Getenv("REDIRECT_URL"),
		ClientID:       os.Getenv("CLIENT_ID"),
		AuthSecret:     os.Getenv("AUTH_SECRET"),
		Oauth2URL:      os.Getenv("OAUTH2_URL"),
		DbUser:         os.Getenv("DB_USER"),
		DbPwd:          os.Getenv("DB_PASSWORD"),
		DbName:         os.Getenv("DB_NAME"),
		TokenSecretJWT: os.Getenv("JWT_TOKEN_SECRET"),
	}
}
