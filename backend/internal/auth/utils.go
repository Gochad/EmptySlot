package auth

import (
	"context"
	"crypto/rand"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

var (
	oauth *oauth2.Config
	Store = sessions.NewCookieStore([]byte(os.Getenv("TOKEN_SECRET")))
)

const (
	tokenSet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	tokenLength = 15
)

type OAuthData struct {
	Id             string `json:"id"`
	Email          string `json:"email"`
	Verified_email bool   `json:"verified_email"`
	Picture        string `json:"picture"`
}

func TokenString() (string, error) {
	charsetLength := len(tokenSet)

	randomBytes := make([]byte, tokenLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	for i := 0; i < tokenLength; i++ {
		randomBytes[i] = tokenSet[int(randomBytes[i])%charsetLength]
	}

	return string(randomBytes), nil
}

func GetUserData(state, code, tokenCode string) ([]byte, error) {
	if state != tokenCode {
		return nil, errors.New("invalid user")
	}

	token, err := oauth.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
