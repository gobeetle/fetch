package authgateway

import (
	"github.com/gobeetle/fetch/internal/token"
)

func (a *Authn) DecodeToken() (*token.TokenClaims, error) {
	return a.base.DecodeToken()
}
