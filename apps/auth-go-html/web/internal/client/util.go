package client

import (
	"encoding/json"
	"net/http"
)

func read(resp *http.Response, v any) *ErrorResponse {
	defer resp.Body.Close()

	var errR ErrorResponse

	if !isSuccess(resp.StatusCode) {
		err := json.NewDecoder(resp.Body).Decode(&errR)
		if err != nil {
			errR.Errors = Errors{"unexpected_decoder_error": err.Error()}
		}
		return &errR
	}

	err := json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		errR.Errors = Errors{"unexpected_decoder_error": err.Error()}
		return &errR
	}
	return nil
}

func readSuccess(resp *http.Response) *ErrorResponse {
	defer resp.Body.Close()

	if !isSuccess(resp.StatusCode) {
		var errR ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&errR)
		if err != nil {
			errR.Errors = Errors{"unexpected_decoder_error": err.Error()}
		}
		return &errR
	}

	return nil
}

func isSuccess(statusCode int) bool {
	return 200 <= statusCode && statusCode <= 299
}
