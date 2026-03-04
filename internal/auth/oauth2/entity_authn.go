package oauth2

import (
	"net/http"
	"strings"
	"time"

	"github.com/gobeetle/fetch/internal/fetch"
	mreq "github.com/gobeetle/fetch/internal/request_modifier"
	mrsp "github.com/gobeetle/fetch/internal/response_modifier"
	"github.com/gobeetle/fetch/internal/token"
	"github.com/gobeetle/fetch/internal/utils"
	"golang.org/x/oauth2"
)

// use private fields with one new function (constructor) to make fields immutable after creation, so that they can't be changed
type Authn struct {
	base          *token.AuthnBase // base struct to manage the token and client
	token_url     string
	client_id     string
	client_secret string // optional as the client might not have a secret, set it to empty string if it is not needed
	username      string // optional for client_credentials, set it to empty string if it is not needed
	password      string // optional for client_credentials, set it to empty string if it is not needed
	scopes        []string
	scope_string  string // space separated scope string to be used in the request
}

/*
GetToken returns the token from the source, only get the token if it is not valid or expired, otherwise return the existing token
*/
func (a *Authn) Token() (*oauth2.Token, error) {
	return a.base.Token()
}

/*
this function always returns a new token from the source
*/
func (a *Authn) NewToken() (*oauth2.Token, error) {
	// otherwise get the token from the server and set it to T as well
	token_resp := TokenResponse{}
	_, err :=
		fetch.New[any, TokenResponse]().
			ModReq(
				mreq.WithMethod[any]("POST"),
				mreq.WithUrl[any](a.token_url),
				mreq.WithForm[any](a.GetOAuthForm()),
			).
			ModRsp(
				mrsp.With2XXAsValidStatusCode[TokenResponse](),
				mrsp.WithJsonObj(&token_resp),
			).Do()

	if err != nil {
		return nil, err
	}
	// convert token_resp to oauth2.Token
	token := oauth2.Token{
		AccessToken:  token_resp.AccessToken,
		TokenType:    token_resp.TokenType,
		RefreshToken: token_resp.RefreshToken,
		Expiry:       time.Now().Add(time.Second * time.Duration(token_resp.ExpiresIn)),
	}

	return &token, nil
}

func (a *Authn) Client() *http.Client {
	return a.base.Client()
}

// GetScopes returns the combined scopes as both a slice and a space-separated string.
// This is useful when you need either format of the scopes.
func (a *Authn) GetScopes() ([]string, string) {
	var scopes []string

	// add scopes from scope_string if not empty
	if a.scope_string != "" {
		scopes = strings.Fields(a.scope_string)
	}

	// add individual scopes if any
	scopes = append(scopes, a.scopes...)

	// remove duplicates if needed
	if len(scopes) > 1 {
		scopes = utils.UniqueStrings(scopes)
	}

	return scopes, strings.Join(scopes, " ")
}

/*
GetOAuthForm returns the form that is used to get the token from the oauth2 server
*/
func (a *Authn) GetOAuthForm() map[string]any {
	form := map[string]any{
		"client_id":     a.client_id,
		"client_secret": a.client_secret,
	}

	// Get the combined scopes
	_, scopeString := a.GetScopes()
	if scopeString != "" {
		form["scope"] = scopeString
	}

	switch a.GetGrantType() {
	case EnumGrantType.PasswordGrant:
		form["grant_type"] = string(EnumGrantType.PasswordGrant)
		form["username"] = a.username
		form["password"] = a.password
	case EnumGrantType.ClientCredentialsGrant:
		form["grant_type"] = string(EnumGrantType.ClientCredentialsGrant)
	}

	return form
}

// GetGrantType determines the OAuth2 grant type based on the provided credentials
func (a *Authn) GetGrantType() GrantType {
	if a == nil {
		return EnumGrantType.UnsupportedGrant
	}
	// if username and password are provided, it's password grant
	if a.username != "" && a.password != "" {
		return EnumGrantType.PasswordGrant
	}
	// if only client credentials are provided, it's client_credentials grant
	if a.client_id != "" && a.client_secret != "" {
		return EnumGrantType.ClientCredentialsGrant
	}
	return EnumGrantType.UnsupportedGrant
}

// ensure that current Authenticator implements the TokenSource interface
var _ token.TokenSource = (*Authn)(nil)
