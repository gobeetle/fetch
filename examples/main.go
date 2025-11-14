package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gobeetle/fetch"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id,omitempty"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	demoGetRequest()
	demoPostRequest()
	demoOAuth2Auth()
}

// demoGetRequest demonstrates how to make a GET request
func demoGetRequest() {
	fmt.Println("=== Example: GET request ===")

	// Create a variable to hold the response
	var post Post

	// Make the GET request
	result, err := fetch.New().
		WithRetry(5, time.Millisecond*100). // with optional retry logic
		ModReq(
			fetch.WithReqUrl("https://jsonplaceholder.typicode.com/posts/1"),
			fetch.WithReqMethod(fetch.EnumMethodType.Get),
		).
		ModRsp(
			fetch.WithRsp2XXAsValidStatusCode(),
			fetch.WithRspJsonObj(&post),
		).Do()

	// Handle errors
	if err != nil {
		log.Printf("GET request failed: %v\n\n", err)
		return
	}

	// Print results
	fmt.Printf("Status: %d\n", result.StatusCode)
	fmt.Printf("Response: %+v\n\n", post)
}

// demoPostRequest demonstrates how to make a POST request
func demoPostRequest() {
	fmt.Println("=== Example: POST request ===")

	// Create a new post to send
	newPost := Post{
		UserId: 1,
		Title:  "Hello, World!",
		Body:   "This is a test post created with the fetch package.",
	}

	// Variable to hold the response
	var createdPost Post

	// Make the POST request
	result, err := fetch.New().
		ModReq(
			fetch.WithReqUrl("https://jsonplaceholder.typicode.com/posts"),
			fetch.WithReqMethod(fetch.EnumMethodType.Post),
			fetch.WithReqJsonBody(newPost),
			fetch.WithReqHeaderContentTypeAsJson(),
		).
		ModRsp(
			fetch.WithRsp2XXAsValidStatusCode(),
			fetch.WithRspJsonObj(&createdPost),
		).Do()

	// Handle errors
	if err != nil {
		log.Printf("POST request failed: %v\n\n", err)
		return
	}

	// Print results
	fmt.Printf("Status: %d\n", result.StatusCode)
	fmt.Printf("Created post: %+v\n\n", createdPost)
}

// demoOAuth2Auth demonstrates how to use fetch.OAuth2Authn for OAuth2 authentication
func demoOAuth2Auth() {
	fmt.Println("=== Example: OAuth2 with fetch.OAuth2Authn ===")

	// Create a new OAuth2 authenticator
	oauth2Authn := fetch.NewOAuth2().
		WithTokenUrl("https://your_token_url_here").
		WithClientId("your_client_id_here").
		WithClientSecret("your_client_secret_here").
		WithUsername("your_username_here").
		WithPassword("your_password_here").
		WithScopes("scope1", "scope2", "scope3"). //add scopes if needed
		WithScopeString("scope4 scope5")          //add scope string if needed

	// Get and print the token
	t, err := oauth2Authn.Token()
	if err != nil {
		log.Printf("Failed to get token: %v\n", err)
		return
	}

	// Print basic token info
	fmt.Printf("Access Token: %s...\n", t.AccessToken[:min(20, len(t.AccessToken))])
	fmt.Printf("Token Type: %s\n", t.TokenType)
	fmt.Printf("Expires In: %v\n", time.Until(t.Expiry).Round(time.Second))

	// this is will get all the associated claims from the token
	// easier for debugging and inspecting the token claims
	claims, _ := oauth2Authn.DecodeToken()
	fmt.Println(claims.TokenClaimsDetails.JsonRawString)
}
