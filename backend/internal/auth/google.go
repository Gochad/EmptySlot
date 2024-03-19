package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	RedirectURL = "REDIRECT_URL"
	ClientID    = "CLIENT_ID"
	Secret      = "AUTH_SECRET"
)

// init loads the neccessary configuration details required by oauth2 package.
func init() {
	oauth = &oauth2.Config{
		RedirectURL:  os.Getenv(RedirectURL),
		ClientID:     os.Getenv(ClientID),
		ClientSecret: os.Getenv(Secret),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func GoogleSignOn(res http.ResponseWriter, req *http.Request) {
	tokenString, err := TokenString()
	if err != nil {
		fmt.Fprintf(res, "error: could not generate random token string: %v", err)
	}

	// creates a new session
	session, err := Store.Get(req, "tokenSession")
	if err != nil {
		fmt.Fprintf(res, "error: %v", err)
	}

	// saves the generated token string into the created session; uses tokenStringKey as the key
	session.Values["tokenStringKey"] = tokenString
	session.Save(req, res)

	// returns a URL with attached tokenString
	url := oauth.AuthCodeURL(tokenString)
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}

func Callback(res http.ResponseWriter, req *http.Request) {
	state := req.FormValue("state")
	code := req.FormValue("code")

	session, err := Store.Get(req, "tokenSession")
	if err != nil {
		fmt.Fprintf(res, "error: %v", err)
	}

	dataToken, ok := session.Values["tokenStringKey"].(string)
	if !ok {
		dataToken = "token not found in the session"
	}

	data, err := GetUserData(state, code, dataToken)
	if err != nil {
		log.Fatal(err)
	}

	session.Options.MaxAge = -1
	session.Save(req, res)

	var authStruct OAuthData

	err = json.Unmarshal([]byte(data), &authStruct)
	if err != nil {
		fmt.Fprintf(res, "error: %v", err)
	}

	status := authStruct.Verified_email
	if status {
		fmt.Fprintf(res, "success: %s is a verified user\n", authStruct.Email)
	} else {
		fmt.Fprint(res, "failed verification")
	}
}
