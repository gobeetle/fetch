// Package testhelper provides utilities for testing HTTP clients and servers.
package testhelper

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

// TestServer is a test HTTP server for testing HTTP clients
type TestServer struct {
	server *httptest.Server
	Mux    *http.ServeMux
}

// NewTestServer creates a new test HTTP server
func NewTestServer() *TestServer {
	mux := http.NewServeMux()
	return &TestServer{
		server: httptest.NewServer(mux),
		Mux:    mux,
	}
}

// URL returns the base URL of the test server
func (ts *TestServer) URL() string {
	return ts.server.URL
}

// Close shuts down the test server
func (ts *TestServer) Close() {
	ts.server.Close()
}

// SetResponse configures the test server to respond to a specific path and method
func (ts *TestServer) SetResponse(method, path string, statusCode int, headers map[string]string, body []byte) {
	ts.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method matches
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Set headers
		for key, value := range headers {
			w.Header().Set(key, value)
		}

		// Set status code and write response body
		w.WriteHeader(statusCode)
		w.Write(body)
	})
}

// AssertRequest verifies that an HTTP request matches the expected values
func AssertRequest(t assert.TestingT, r *http.Request, method, path string, body []byte) bool {
	if !assert.Equal(t, method, r.Method, "Unexpected request method") {
		return false
	}

	if !assert.Equal(t, path, r.URL.Path, "Unexpected request path") {
		return false
	}

	if body != nil {
		defer r.Body.Close()
		actualBody, err := io.ReadAll(r.Body)
		if !assert.NoError(t, err, "Failed to read request body") {
			return false
		}
		if !assert.Equal(t, body, actualBody, "Unexpected request body") {
			return false
		}
	}

	return true
}
