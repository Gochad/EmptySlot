package internal

import (
	"os"
)

type oauth struct {
	TokenSecret string
	RedirectURL string
	ClientID    string
	AuthSecret  string
	URL         string
}

type session struct {
	DbUser         string
	DbPwd          string
	DbName         string
	TokenSecretJWT string
}

type config struct {
	Oauth   oauth
	Session session
}

var EnvConfig config

func init() {
	EnvConfig = config{
		Oauth: oauth{
			TokenSecret: os.Getenv("TOKEN_SECRET"),
			RedirectURL: os.Getenv("REDIRECT_URL"),
			ClientID:    os.Getenv("CLIENT_ID"),
			AuthSecret:  os.Getenv("AUTH_SECRET"),
			URL:         os.Getenv("OAUTH2_URL"),
		},
		Session: session{
			DbUser:         os.Getenv("DB_USER"),
			DbPwd:          os.Getenv("DB_PASSWORD"),
			DbName:         os.Getenv("DB_NAME"),
			TokenSecretJWT: os.Getenv("JWT_TOKEN_SECRET"),
		},
	}
}
