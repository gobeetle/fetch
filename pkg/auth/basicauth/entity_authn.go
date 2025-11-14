package basicauth

import (
	"encoding/base64"
	"net/http"

	"github.com/gobeetle/fetch"
	"golang.org/x/oauth2"
)

// Authn implements basic authentication
// It uses private fields with a constructor to make fields immutable after creation
type Authn struct {
	username string
	password string
	base     *fetch.AuthnBase // base struct to manage the token and client
}

// Token returns the token from the source, only get the token if it is not valid or expired, otherwise return the existing token
func (a *Authn) Token() (*oauth2.Token, error) {
	return a.base.Token()
}

// NewToken always returns a new basic auth token
// For basic auth, the token is just a base64 encoded string of "username:password"
func (a *Authn) NewToken() (*oauth2.Token, error) {
	token := oauth2.Token{
		TokenType:   "Basic",
		AccessToken: a.getBasicAuthToken(),
	}
	return &token, nil
}

// Client returns an HTTP client that adds basic auth to requests
func (a *Authn) Client() *http.Client {
	return a.base.Client()
}

// getBasicAuthToken returns the base64 encoded "username:password" string
func (a *Authn) getBasicAuthToken() string {
	auth := a.username + ":" + a.password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// ensure that Authn implements the TokenSource interface
var _ fetch.TokenSource = (*Authn)(nil)
