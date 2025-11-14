package oauth2

type GrantType string

var EnumGrantType = struct {
	ClientCredentialsGrant GrantType
	PasswordGrant          GrantType
	UnsupportedGrant       GrantType
}{
	ClientCredentialsGrant: "client_credentials",
	PasswordGrant:          "password",
	UnsupportedGrant:       "unsupported",
}
