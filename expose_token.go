package fetch

import "github.com/gobeetle/fetch/internal/token"

type (
	AuthnBase          = token.AuthnBase
	TokenCallback      = token.TokenCallback
	TokenSource        = token.TokenSource
	TokenClaims        = token.TokenClaims
	TokenClaimsDetails = token.TokenClaimsDetails
)

const ()

var (
	NewAuthnBase = token.NewAuthnBase
)
