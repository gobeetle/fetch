package fetch

import (
	"github.com/gobeetle/fetch/internal/token"
	"golang.org/x/oauth2"
)

type (
	AuthnBase          = token.AuthnBase
	TokenCallback      = token.TokenCallback
	TokenSource        = token.TokenSource
	TokenClaims        = token.TokenClaims
	TokenClaimsDetails = token.TokenClaimsDetails
)

const ()

var ()

func NewAuthnBase(tokenObtainer func() (*oauth2.Token, error)) *AuthnBase {
	return token.NewAuthnBase(tokenObtainer)
}
