package token

import (
	"net/http"

	"golang.org/x/oauth2"
)

/*
inspired by oauth2.TokenSource, this interface will be used to get the token from the source and also get a new token
implement this interface together with AuthenticatorBase to get the token from the source
AuthenticatorBase manage the token and client, it will also have a callback function to be called when a new token is obtained
see fauth modules for more examples
*/
type TokenSource interface {
	oauth2.TokenSource // this will get a valid token, cached valid token will be used if it is not expired, otherwise it will get a new token
	Validate() error
	NewToken() (*oauth2.Token, error) // get a new token, this will enforce the token to be refreshed
	Client() *http.Client             // get a http client
	DecodeToken() (*TokenClaims, error)
}
