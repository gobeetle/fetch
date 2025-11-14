package token

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenClaims represents the standard JWT claims plus any custom claims
type TokenClaims struct {
	jwt.RegisteredClaims
	TokenClaimsDetails
	UserID        string    `json:"sub,omitempty"`
	Email         string    `json:"email,omitempty"`
	Scope         string    `json:"scope,omitempty"`
	Scopes        []string  `json:"scopes,omitempty"` // Alternative if scope is provided as array
	ClientID      string    `json:"client_id,omitempty"`
	IssuedAt      int64     `json:"iat,omitempty"`      // Issued At (timestamp)
	IssuedAtTime  time.Time `json:"iat_time,omitempty"` // Issued At (timestamp)
	ExpiresAt     int64     `json:"exp,omitempty"`      // Expiration Time (timestamp)
	ExpiresAtTime time.Time `json:"exp_time,omitempty"` // Expiration Time (timestamp)
	NotBefore     int64     `json:"nbf,omitempty"`      // Not Before (timestamp)
	NotBeforeTime time.Time `json:"nbf_time,omitempty"` // Not Before (timestamp)
	Issuer        string    `json:"iss,omitempty"`      // Issuer
	Audience      []string  `json:"aud,omitempty"`      // Audience
	JWTID         string    `json:"jti,omitempty"`      // JWT ID
}

type TokenClaimsDetails struct {
	JsonRawMessage json.RawMessage    `json:"claims_json_message,omitempty"`
	JsonRawString  string             `json:"claims_json_string,omitempty"`
	AnyMap         map[string]any     `json:"claims_any_map,omitempty"`
	Int32Map       map[string]int32   `json:"claims_int32_map,omitempty"`
	Int64Map       map[string]int64   `json:"claims_int64_map,omitempty"`
	Float64Map     map[string]float64 `json:"claims_float64_map,omitempty"`
	BoolMap        map[string]bool    `json:"claims_bool_map,omitempty"`
	StringMap      map[string]string  `json:"claims_string_map,omitempty"`
}

// Valid implements jwt.Claims interface
// Note: The actual validation is handled by the jwt library internally
// when parsing with Parse() instead of ParseUnverified()

// GetExpirationTime returns the expiration time of the token
func (c *TokenClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	if c.ExpiresAt == 0 {
		return nil, nil
	}
	return jwt.NewNumericDate(time.Unix(c.ExpiresAt, 0)), nil
}

// GetIssuedAt returns the issued at time of the token
func (c *TokenClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	if c.IssuedAt == 0 {
		return nil, nil
	}
	return jwt.NewNumericDate(time.Unix(c.IssuedAt, 0)), nil
}

// GetNotBefore returns the not before time of the token
func (c *TokenClaims) GetNotBefore() (*jwt.NumericDate, error) {
	if c.NotBefore == 0 {
		return nil, nil
	}
	return jwt.NewNumericDate(time.Unix(c.NotBefore, 0)), nil
}

// GetIssuer returns the issuer of the token
func (c *TokenClaims) GetIssuer() (string, error) {
	return c.Issuer, nil
}

// GetSubject returns the subject of the token (usually user ID)
func (c *TokenClaims) GetSubject() (string, error) {
	return c.UserID, nil
}

// GetAudience returns the audience of the token
func (c *TokenClaims) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings(c.Audience), nil
}

var _ jwt.Claims = (*TokenClaims)(nil)
