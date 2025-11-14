package authgateway

/*
Response is a struct that represents the response (wrapper) from the auth gate way service
*/
type Response struct {
	Status  int `json:"status"`
	Data    any
	Message string `json:"message"`
}

/*
TokenResponse is a struct that represents the response from the auth gate way service
*/
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Name         string `json:"name"`
	Email        string `json:"email"`
}
