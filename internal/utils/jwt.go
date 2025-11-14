package utils

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

// DecodeJWT parses a JWT token and returns its claims as a map.
// If the token is invalid or malformed, it returns an error.
func DecodeJWT(tokenString string) (jwt.MapClaims, error) {
	// Remove 'Bearer ' prefix if present
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse the token without verification since we just want to read the claims
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	// Type assert the claims to MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims format")
	}

	return claims, nil
}

// PrettyPrintJWT prints the JWT claims in a human-readable format.
func PrettyPrintJWT(tokenString string) (string, error) {
	claims, err := DecodeJWT(tokenString)
	if err != nil {
		return "", err
	}

	var result strings.Builder
	result.WriteString("JWT Claims:\n")
	for k, v := range claims {
		switch v := v.(type) {
		case float64:
			// Handle JSON numbers that might be represented as float64
			if v == float64(int64(v)) {
				fmt.Fprintf(&result, "  %s: %d\n", k, int64(v))
			} else {
				fmt.Fprintf(&result, "  %s: %f\n", k, v)
			}
		case map[string]interface{}:
			// Handle nested objects
			fmt.Fprintf(&result, "  %s: %+v\n", k, v)
		case []interface{}:
			// Handle arrays
			fmt.Fprintf(&result, "  %s: %v\n", k, v)
		default:
			// Handle strings and other types
			fmt.Fprintf(&result, "  %s: %v\n", k, v)
		}
	}

	return result.String(), nil
}

// DecodeOAuth2Token decodes an oauth2.Token's access token and returns the claims as a formatted string.
func DecodeOAuth2Token(token *oauth2.Token) (string, error) {
	if token == nil || token.AccessToken == "" {
		return "", fmt.Errorf("invalid or empty token")
	}
	return PrettyPrintJWT(token.AccessToken)
}
