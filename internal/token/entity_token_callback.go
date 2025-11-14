package token

import "golang.org/x/oauth2"

type TokenCallback func(*oauth2.Token) error
