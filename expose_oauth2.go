package fetch

import "github.com/gobeetle/fetch/internal/auth/oauth2"

type (
	// OAuth2 represents an OAuth2 authenticator
	OAuth2Authn = oauth2.Authn

	// OAuth2TokenResponse represents the token response from OAuth2 server
	OAuth2TokenResponse = oauth2.TokenResponse

	// OAuth2GrantType represents the OAuth2 grant type
	OAuth2GrantType = oauth2.GrantType
)

const ()

var (
	// EnumOAuth2GrantType provides the available OAuth2 grant types
	EnumOAuth2GrantType = oauth2.EnumGrantType
)

// NewOAuth2 creates a new OAuth2 authenticator
func NewOAuth2() *OAuth2Authn {
	return oauth2.New()
}
