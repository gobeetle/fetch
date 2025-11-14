# Fetch & Authenticator Package
![Coverage](https://img.shields.io/badge/Coverage-74.0%25-brightgreen)

This package provides a set of tools for managing HTTP requests and authentication in a Go application. It includes the following core components:

1. **Fetch**: A flexible HTTP request utility that supports various options such as retry logic, request/response modifiers, and error handling.
2. **Authenticator**: A utility for obtaining and managing OAuth2 tokens from an authentication gateway, with automatic token refreshing.

## Features

- **HTTP Request Handling**: Simplifies making HTTP requests with built-in support for setting request options (e.g., method, headers, body) and response processing.
- **Retry Logic**: Automatically retries failed requests with configurable retry count and delay.
- **Token Management**: Facilitates obtaining and refreshing OAuth2 tokens, integrating seamlessly with HTTP clients.

## Installation

To install the package, run:


go get github.com/gobeetle/fetch



---
## Usage
#### Fetch
The Fetch is used to make HTTP requests with various options and automatic retry logic.

Example: Basic GET Request

```go
package main

import (
    "fmt"
    "github.com/gobeetle/fetch"
)

func main() {
    result, err := fetch.New().
        SetReq(
            fetch.SetMethod("GET"), 
            fetch.SetUrl("http://example.com")
        ).Do()

    if err != nil {
        fmt.Println("Request failed:", err)
        return
    }

    fmt.Println("Response:", string(result.RespBytes))
}
```

Example: POST Request with JSON Body

```go
package main

import (
    "fmt"
    "github.com/gobeetle/fetch"
)

func main() {
    data := map[string]string{"key": "value"}
    result, err := fetch.New().
        SetReq(
            fetch.SetMethod("POST"),
            fetch.SetUrl("http://example.com/api"),
            fetch.SetJsonBody(data),
        ).
        Do()

    if err != nil {
        fmt.Println("Request failed:", err)
        return
    }

    fmt.Println("Response:", string(result.RespBytes))
}
```

---

#### Authenticator

The Authenticator is used to manage OAuth2 tokens, automatically obtaining and refreshing them as needed.

Example: Obtaining a Token

```go
package main

import (
    "fmt"
    "github.com/gobeetle/fetch/fauth/authgateway"
)

func main() {
    auth := authgateway.New("http://example.com/login?key=%s", "your_access_key", nil)
    token, err := auth.GetToken()

    if err != nil {
        fmt.Println("Failed to get token:", err)
        return
    }

    fmt.Println("Access Token:", token.AccessToken)
}
```

Example: Using Authenticator with an HTTP Client

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/gobeetle/fetch/fauth/authgateway"
)

func main() {
    auth := authgateway.New("http://example.com/login?key=%s", "your_access_key", nil)
    client := auth.Client()

    req, _ := http.NewRequest("GET", "http://example.com/protected", nil)
    resp, err := client.Do(req)

    if err != nil {
        fmt.Println("Request failed:", err)
        return
    }

    defer resp.Body.Close()
    fmt.Println("Response Status:", resp.Status)
}

```

---

#### Error Handling
The package provides a custom error type FError for detailed error handling, including status codes and error wrapping.

Example: Handling Errors

```go
package main

import (
    "fmt"
    "github.com/gobeetle/fetch"
)

func main() {
    _, err := fetch.New().
        SetReq(fetch.SetMethod("GET"), fetch.SetUrl("http://invalid-url")).
        Do()

    if err != nil {
        if fErr, ok := err.(fetch.FError); ok {
            fmt.Printf("Request failed with status code %d: %s\n", fErr.StatusCode(), fErr.Error())
        } else {
            fmt.Println("Request failed:", err)
        }
    }
}
```

#### Contributing
Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.