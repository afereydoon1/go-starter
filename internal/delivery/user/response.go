package user

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error string `json:"error" example:"Error message"`
}

// TokenResponse represents a successful login response
type TokenResponse struct {
	Token string `json:"token" example:"jwt-token"`
}
