package oauth2

import (
	"github.com/gobeetle/fetch/internal/token"
)

// DecodeToken decodes the access token and returns the claims as a TokenClaims struct
// The raw token string is stored in the ClaimsString field of the returned struct.
func (a *Authn) DecodeToken() (*token.TokenClaims, error) {
	return a.base.DecodeToken()
}
