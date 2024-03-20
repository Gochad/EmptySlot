package auth

import (
	"context"
	"errors"
	"io"
	mr "math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"

	"backend/internal"
)

var (
	oauth *oauth2.Config
	Store = sessions.NewCookieStore([]byte(internal.EnvConfig.TokenSecret))
)

type OAuthData struct {
	Id             string `json:"id"`
	Email          string `json:"email"`
	Verified_email bool   `json:"verified_email"`
	Picture        string `json:"picture"`
}

const (
	tokenSet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	tokenLength = 15
)

func init() {
	mr.New(mr.NewSource(time.Now().UnixNano()))
}

func TokenString() string {
	var sb strings.Builder
	sb.Grow(tokenLength)

	for i := 0; i < tokenLength; i++ {
		randomIndex := mr.Intn(len(tokenSet))
		sb.WriteByte(tokenSet[randomIndex])
	}

	return sb.String()
}

func GetUserData(state, code, tokenCode string) ([]byte, error) {
	if state != tokenCode {
		return nil, errors.New("invalid user")
	}

	token, err := oauth.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	response, err := http.Get(internal.EnvConfig.Oauth2URL + token.AccessToken)
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
