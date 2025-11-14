package token

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/oauth2"
)

// AuthnBase manages OAuth2 tokens and provides an HTTP client that automatically handles token refresh.
// It implements the oauth2.TokenSource interface and is safe for concurrent use.
type AuthnBase struct {
	// mu               sync.RWMutex
	clientMutex      sync.Mutex
	token            *oauth2.Token
	client           *http.Client
	onTokenRefreshed TokenCallback
	newTokenObtainer func() (*oauth2.Token, error)
	closeOnce        sync.Once
	closed           bool
}

// NewAuthnBase creates a new AuthnBase with the provided token obtainer function.
// The tokenObtainer function will be called when a new token is needed.
func NewAuthnBase(tokenObtainer func() (*oauth2.Token, error)) *AuthnBase {
	if tokenObtainer == nil {
		panic("tokenObtainer cannot be nil")
	}
	return &AuthnBase{
		newTokenObtainer: tokenObtainer,
	}
}

// WithOnTokenRefreshed sets a callback that will be called whenever a new token is obtained.
// The callback is called with the new token and any error that occurred during refresh.
func (a *AuthnBase) WithOnTokenRefreshed(callback TokenCallback) *AuthnBase {
	if a == nil {
		return nil
	}

	if a.closed {
		return a
	}

	a.onTokenRefreshed = callback
	return a
}

// Token returns a valid token, refreshing it if necessary.
// Implements the oauth2.TokenSource interface.
func (a *AuthnBase) Token() (*oauth2.Token, error) {
	// Fast path: check if we have a valid token with read lock

	if a.token != nil && a.token.Valid() && !a.closed {
		token := a.token
		return token, nil
	}

	// Slow path: need to get a new token
	return a.NewToken()
}

func (a *AuthnBase) NewToken() (*oauth2.Token, error) {
	// Double-check after acquiring write lock
	if a.closed {
		return nil, fmt.Errorf("AuthnBase is closed")
	}

	// Check again in case another goroutine refreshed the token
	if a.token != nil && a.token.Valid() {
		return a.token, nil
	}

	// Get a new token
	newToken, err := a.newTokenObtainer()
	if err != nil {
		return nil, fmt.Errorf("failed to obtain new token: %w", err)
	}

	// Update the token
	a.token = newToken

	// Call the callback if set
	if a.onTokenRefreshed != nil {
		if err := a.onTokenRefreshed(newToken); err != nil {
			return nil, fmt.Errorf("token refresh callback failed: %w", err)
		}
	}

	return newToken, nil
}

// Client returns an HTTP client that automatically handles token refresh.
// The same client instance is returned on subsequent calls.
// If the AuthnBase is closed, it will return nil.
func (a *AuthnBase) Client() *http.Client {
	a.clientMutex.Lock()
	defer a.clientMutex.Unlock()

	if a.client != nil {
		return a.client
	}

	if a.closed {
		return nil
	}

	a.client = &http.Client{
		Transport: &oauth2.Transport{
			Source: a,
			Base:   http.DefaultTransport,
		},
		Timeout: 30 * time.Second,
	}

	return a.client
}

// Close releases any resources used by the AuthnBase.
// After calling Close, the AuthnBase should not be used.
func (a *AuthnBase) Close() error {
	if a == nil {
		return nil
	}

	var err error
	a.closeOnce.Do(func() {
		a.closed = true
		if a.client != nil {
			// If the client has a custom transport with a connection pool, close it
			if transport, ok := a.client.Transport.(*http.Transport); ok {
				transport.CloseIdleConnections()
			}
			a.client = nil
		}
		a.token = nil
	})

	return err
}

// Valid returns true if the AuthnBase has a valid token and is not closed.
func (a *AuthnBase) Valid() bool {
	if a == nil {
		return false
	}

	return !a.closed && a.token != nil && a.token.Valid()
}

// DecodeToken decodes the access token and returns the claims as a TokenClaims struct
// The raw token string is stored in the ClaimsString field of the returned struct.
func (a *AuthnBase) DecodeToken() (*TokenClaims, error) {
	t, err := a.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	return DecodeToken(t)
}

// Validate returns an error if the token is invalid or expired.
func (a *AuthnBase) Validate() error {
	if a == nil {
		return nil
	}
	return nil
}

var _ TokenSource = (*AuthnBase)(nil)
