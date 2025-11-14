package oauth2

import (
	"github.com/gobeetle/fetch/internal/token"
)

func New() *Authn {
	authn := &Authn{}
	authn.base = token.NewAuthnBase(authn.NewToken)
	return authn
}

func (a *Authn) WithOnTokenRefreshed(on_token_refresh token.TokenCallback) *Authn {
	if a == nil {
		return nil
	}
	a.base = a.base.WithOnTokenRefreshed(on_token_refresh)
	return a
}

func (a *Authn) WithTokenUrl(token_url string) *Authn {
	if a == nil {
		return nil
	}
	a.token_url = token_url
	return a
}

func (a *Authn) WithClientId(client_id string) *Authn {
	if a == nil {
		return nil
	}
	a.client_id = client_id
	return a
}

func (a *Authn) WithClientSecret(client_secret string) *Authn {
	if a == nil {
		return nil
	}
	a.client_secret = client_secret
	return a
}

func (a *Authn) WithScopes(scopes ...string) *Authn {
	if a == nil {
		return nil
	}
	a.scopes = scopes
	return a
}

func (a *Authn) WithScopeString(scope string) *Authn {
	if a == nil {
		return nil
	}
	a.scope_string = scope
	return a
}

func (a *Authn) WithUsername(username string) *Authn {
	if a == nil {
		return nil
	}
	a.username = username
	return a
}

func (a *Authn) WithPassword(password string) *Authn {
	if a == nil {
		return nil
	}
	a.password = password
	return a
}
