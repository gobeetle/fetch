package token

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gobeetle/fetch/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

func DecodeTokenSource(tkn oauth2.TokenSource) (*TokenClaims, error) {
	if tkn == nil {
		return nil, fmt.Errorf("token source cannot be nil")
	}
	token, err := tkn.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	return DecodeToken(token)
}

func DecodeTokenString(tokenString string) (*TokenClaims, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("token string cannot be empty")
	}
	return DecodeToken(&oauth2.Token{
		AccessToken: tokenString,
	})
}

// DecodeToken decodes the access token and returns the claims as a TokenClaims struct
// The raw token string is stored in the ClaimsString field of the returned struct.
func DecodeToken(tkn *oauth2.Token) (*TokenClaims, error) {
	if tkn == nil {
		return nil, fmt.Errorf("token is nil")
	}

	// Remove 'Bearer ' prefix if present
	tokenString := strings.TrimPrefix(tkn.AccessToken, "Bearer ")

	// Parse the token into a map to extract all claims
	parser := new(jwt.Parser)
	tokenObj, _, err := parser.ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Get the claims as a map
	mapClaims, ok := tokenObj.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims format")
	}

	// Create our claims struct
	claims := &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	// Manually map the standard claims
	if exp, ok := mapClaims["exp"].(float64); ok {
		expTime := time.Unix(int64(exp), 0)
		claims.ExpiresAt = int64(exp)
		claims.ExpiresAtTime = expTime
	}
	if iat, ok := mapClaims["iat"].(float64); ok {
		iatTime := time.Unix(int64(iat), 0)
		claims.IssuedAt = int64(iat)
		claims.IssuedAtTime = iatTime
	}
	if nbf, ok := mapClaims["nbf"].(float64); ok {
		nbfTime := time.Unix(int64(nbf), 0)
		claims.NotBefore = int64(nbf)
		claims.NotBeforeTime = nbfTime
	}
	if iss, ok := mapClaims["iss"].(string); ok {
		claims.Issuer = iss
	}
	if sub, ok := mapClaims["sub"].(string); ok {
		claims.UserID = sub
	}
	if aud, ok := mapClaims["aud"]; ok {
		switch v := aud.(type) {
		case string:
			claims.Audience = []string{v}
		case []interface{}:
			audience := make([]string, 0, len(v))
			for _, a := range v {
				if s, ok := a.(string); ok {
					audience = append(audience, s)
				}
			}
			claims.Audience = audience
		}
	}
	if jti, ok := mapClaims["jti"].(string); ok {
		claims.JWTID = jti
	}

	// Handle custom claims
	if scope, ok := mapClaims["scope"].(string); ok && scope != "" {
		claims.Scope = scope
		claims.Scopes = strings.Fields(scope)
	}
	if email, ok := mapClaims["email"].(string); ok && email != "" {
		claims.Email = email
	}
	if clientID, ok := mapClaims["client_id"].(string); ok && clientID != "" {
		claims.ClientID = clientID
	}

	// Convert claims to JSON string for JsonString
	claimsJSON, err := json.Marshal(mapClaims)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal claims: %w", err)
	}

	// Initialize ClaimsDetails
	details := TokenClaimsDetails{
		JsonRawMessage: claimsJSON,
		JsonRawString:  string(claimsJSON),
		AnyMap:         mapClaims,
	}

	// Categorize claims into type-specific maps using the utils package
	int32Map, int64Map, float64Map, boolMap, stringMap := utils.CategorizeMap(mapClaims)
	details.Int32Map = int32Map
	details.Int64Map = int64Map
	details.Float64Map = float64Map
	details.BoolMap = boolMap
	details.StringMap = stringMap

	// Assign the populated details to the claims
	claims.TokenClaimsDetails = details

	return claims, nil
}
