package basicauth

import "github.com/gobeetle/fetch"

// New creates a new basic auth authenticator
func New() *Authn {
	authn := &Authn{}
	authn.base = fetch.NewAuthnBase(authn.NewToken)
	return authn
}

// WithOnTokenRefreshed sets the callback function that will be called when a new token is obtained
func (a *Authn) WithOnTokenRefreshed(onTokenRefresh fetch.TokenCallback) *Authn {
	if a == nil {
		return nil
	}
	a.base = a.base.WithOnTokenRefreshed(onTokenRefresh)
	return a
}

// WithUsername sets the username for basic authentication
func (a *Authn) WithUsername(username string) *Authn {
	if a == nil {
		return nil
	}
	a.username = username
	return a
}

// WithPassword sets the password for basic authentication
func (a *Authn) WithPassword(password string) *Authn {
	if a == nil {
		return nil
	}
	a.password = password
	return a
}
