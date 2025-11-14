package basicauth

import "errors"

// Validate checks if the basic auth configuration is valid
// It implements the TokenSource interface
func (a *Authn) Validate() error {
	if a == nil {
		return errors.New("basic auth cannot be nil")
	}
	if a.username == "" {
		return errors.New("username cannot be empty for basic auth")
	}
	if a.password == "" {
		return errors.New("password cannot be empty for basic auth")
	}
	return nil
}
