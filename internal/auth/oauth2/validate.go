package oauth2

import (
	"errors"
	"strings"
)

func (a *Authn) Validate() error {
	if a == nil {
		return errors.New("oauth2 authn cannot be nil")
	}

	if len(strings.TrimSpace(a.token_url)) == 0 {
		return errors.New("token_url cannot be empty for oauth2 authn")
	}

	if len(strings.TrimSpace(a.client_id)) == 0 {
		return errors.New("client_id cannot be empty for oauth2 authn")
	}
	switch a.GetGrantType() {
	case EnumGrantType.PasswordGrant:
		// for password grant, we've already validated that both username and password are provided
		return nil
	case EnumGrantType.ClientCredentialsGrant:
		// for client credentials, we just need client_id and client_secret
		return nil
	default:
		return errors.New("invalid oauth2 configuration: missing required credentials for any supported grant type")
	}
}
