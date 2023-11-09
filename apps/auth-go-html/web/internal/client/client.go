package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// APIClient is the client helper to interact with the API
type APIClient struct {
	baseURL    string
	httpClient *http.Client
}

// AuthResponse represents the response from the authenticate endpoint of the API
type AuthResponse struct {
	Id string `json:"id"`
}

// Profile represents a user model that matches the one we have in the API
type Profile struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	External  bool   `json:"external"`
	Errors    Errors `json:"errors"`
}

type SignUpResponse struct {
	Profile Profile `json:"profile"`
}

type ProfileResponse struct {
	Profile Profile `json:"profile"`
}

type PatchProfile struct {
	Email     string `json:"email,omitempty"`
	Name      string `json:"name,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	External  string `json:"external,omitempty"`
	Errors    Errors `json:"errors,omitempty"`
}

type Errors map[string]string

// ErrorResponse matches the standard response from the API when errors occur
type ErrorResponse struct {
	Errors Errors `json:"errors"`
}

// NewAPIClient creates a new APIClient
func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

// Authenticate authenticates a user with the API and returns the userid
func (c *APIClient) Authenticate(email, password string) (string, Errors) {
	data := map[string]string{
		"email":    email,
		"password": password,
	}
	jsonData, _ := json.Marshal(data)

	resp, err := c.httpClient.Post(fmt.Sprintf("%s/login", c.baseURL), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", Errors{"unexpected_request_error": err.Error()}
	}

	var response AuthResponse
	errR := read(resp, &response)
	if errR != nil {
		return "", errR.Errors
	}

	return response.Id, nil
}

func (c *APIClient) SignUp(email, password string) *Profile {
	data := map[string]string{
		"email":    email,
		"password": password,
	}
	jsonData, _ := json.Marshal(data)

	resp, err := c.httpClient.Post(fmt.Sprintf("%s/signup", c.baseURL), "application/json", bytes.NewBuffer(jsonData))
	var response SignUpResponse
	if err != nil {
		response.Profile.Email = email
		response.Profile.Errors = Errors{"unexpected_request_error": err.Error()}
		return &response.Profile
	}

	errR := read(resp, &response)
	if errR != nil {
		response.Profile.Errors = errR.Errors
	}

	return &response.Profile
}

// GetProfile fetches the profile information of a user from the API
func (c *APIClient) GetProfile(token string) *Profile {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/profile", c.baseURL), nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := c.httpClient.Do(req)
	var response ProfileResponse
	if err != nil {
		response.Profile.Errors = Errors{"unexpected_request_error": err.Error()}
		return &response.Profile
	}

	errR := read(resp, &response)
	if errR != nil {
		response.Profile.Errors = errR.Errors
	}

	return &response.Profile
}

// UpdateProfile updates the profile information of a user in the API
func (c *APIClient) UpdateProfile(profile *PatchProfile, token string) Errors {
	jsonData, _ := json.Marshal(profile)
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("%s/profile", c.baseURL), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Errors{"unexpected_request_error": err.Error()}
	}

	errR := readSuccess(resp)
	if errR != nil {
		return errR.Errors
	}

	return nil
}
