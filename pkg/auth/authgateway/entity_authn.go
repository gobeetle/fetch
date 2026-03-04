package authgateway

/*
package fauth provides a way to get token from auth gateway service
*/

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gobeetle/fetch"
	"golang.org/x/oauth2"
	"gopkg.in/square/go-jose.v2/jwt"
)

// use private fields with one new function (constructor) to make fields immutable after creation, so that they can't be changed
type Authn struct {
	login_url_format string
	access_key       string
	base             *fetch.AuthnBase // base struct to manage the token and client
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
	resp := Response{Data: &TokenResponse{}}
	_, err :=
		fetch.New[any, Response]().
			ModReq(
				fetch.WithReqMethod[any]("GET"),
				fetch.WithReqUrl[any](fmt.Sprintf(a.login_url_format, a.access_key)),
			).
			ModRsp(
				fetch.WithRsp2XXAsValidStatusCode[Response](),
				fetch.WithRspJsonObj[Response](&resp),
			).Do()

	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, fmt.Errorf("no data in the token response")
	}

	token_resp, ok := resp.Data.(*TokenResponse)
	if !ok {
		return nil, fmt.Errorf("invalid token response")
	}

	parsedJwt, parseErr := jwt.ParseSigned(token_resp.AccessToken)
	if parseErr != nil {
		return nil, parseErr
	}

	var claims map[string]any
	if err := parsedJwt.UnsafeClaimsWithoutVerification(&claims); err != nil {
		return nil, err
	}

	token_exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("exp not found in the token")
	}

	token_type, ok := claims["typ"].(string)
	if !ok {
		return nil, fmt.Errorf("typ not found in the token")
	}

	// convert token_resp to oauth2.Token
	token := oauth2.Token{
		AccessToken:  token_resp.AccessToken,
		TokenType:    token_type,
		RefreshToken: token_resp.RefreshToken,
		Expiry:       time.Unix(int64(token_exp), 0),
	}

	return &token, nil
}

func (a *Authn) Client() *http.Client {
	return a.base.Client()
}

// ensure that current Authenticator implements the TokenSource interface
var _ fetch.TokenSource = (*Authn)(nil)
