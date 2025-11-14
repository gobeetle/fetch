# Fetch & Authenticator Package

This package provides a set of tools for managing HTTP requests and authentication in a Go application. It includes the following core components:

1. **Fetch**: A flexible HTTP request utility that supports various options such as retry logic, request/response modifiers, and error handling.
2. **Authenticator**: A utility for obtaining and managing OAuth2 tokens from an authentication gateway, with automatic token refreshing.

## Features

- **HTTP Request Handling**: Simplifies making HTTP requests with built-in support for setting request options (e.g., method, headers, body) and response processing.
- **Retry Logic**: Automatically retries failed requests with configurable retry count and delay.
- **Token Management**: Facilitates obtaining and refreshing OAuth2 tokens, integrating seamlessly with HTTP clients.

## Installation

To install the package, run:

```
go get github.com/gobeetle/fetch
```


---
## Examples

You can run the live examples from this repository using the Go toolchain.

```bash
go run ./examples
```

This will run the `examples/main.go` program, which demonstrates:

- **GET request**: Fetches a post from `https://jsonplaceholder.typicode.com/posts/1` and prints the decoded JSON as a `Post` struct.
- **POST request**: Sends a JSON payload to `https://jsonplaceholder.typicode.com/posts` and prints the created resource returned by the API.
- **OAuth2 authentication**: Shows how to configure `fetch.NewOAuth2()` with token URL, client credentials, username/password, and scopes, then retrieves and inspects the access token and its claims.

Update the URLs and credentials in `examples/main.go` as needed to match your environment when trying out OAuth2.

---
## Usage

> For the most accurate and up-to-date examples, use the `examples` folder as the source of truth.

- **Fetch**
  - High-level helper around `net/http`.
  - Supports request/response modifiers, retry logic, JSON handling, and status-code filtering.

- **OAuth2 Authenticator**
  - Provides `fetch.NewOAuth2()` for obtaining and refreshing tokens.
  - Can also produce an `*http.Client` that automatically injects and refreshes tokens.

- **Error Handling**
  - Uses a custom `fetch.Error` type that can wrap multiple errors and carry status code, response body, and messages.
