package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"backend/internal"
)

func init() {
	oauth = &oauth2.Config{
		RedirectURL:  internal.EnvConfig.Oauth.RedirectURL,
		ClientID:     internal.EnvConfig.Oauth.ClientID,
		ClientSecret: internal.EnvConfig.Oauth.AuthSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func GoogleSignOn(res http.ResponseWriter, req *http.Request) {
	tokenString := TokenString()

	session, err := Store.Get(req, "tokenSession")
	if err != nil {
		fmt.Fprintf(res, "error: %v", err)
	}

	session.Values["tokenStringKey"] = tokenString
	session.Save(req, res)

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

	err = json.Unmarshal(data, &authStruct)
	if err != nil {
		fmt.Fprintf(res, "error: %v", err)
	}

	status := authStruct.Verified_email
	if status {
		http.Redirect(res, req, "http://localhost:3001/dashboard", http.StatusFound)
		fmt.Fprintf(res, "success: %s is a verified user\n", authStruct.Email)
	} else {
		http.Redirect(res, req, "http://localhost:3001/login", http.StatusFound)
		fmt.Fprint(res, "failed verification")
	}
}
