package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/people/v1"
)

// User info stored in context
type ContextUser struct {
    ID    int
    Email string
}

var Oauth *oauth2.Config;

func InitOauthConfig() {
	Oauth = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile", 
			"https://www.googleapis.com/auth/userinfo.email",
			gmail.GmailReadonlyScope,
			people.ContactsReadonlyScope,
			people.DirectoryReadonlyScope,
		},
		Endpoint: google.Endpoint,
	}
}

