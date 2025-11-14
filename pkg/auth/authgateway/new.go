package authgateway

import "github.com/gobeetle/fetch"

func New() *Authn {
	authn := &Authn{}
	authn.base = fetch.NewAuthnBase(authn.NewToken)
	return authn
}

func (a *Authn) WithOnTokenRefreshed(on_token_refresh fetch.TokenCallback) *Authn {
	if a == nil {
		return nil
	}
	a.base = a.base.WithOnTokenRefreshed(on_token_refresh)
	return a
}

func (a *Authn) WithLoginUrlFormat(login_url_format string) *Authn {
	if a == nil {
		return nil
	}
	a.login_url_format = login_url_format
	return a
}

func (a *Authn) WithAccessKey(access_key string) *Authn {
	if a == nil {
		return nil
	}
	a.access_key = access_key
	return a
}
