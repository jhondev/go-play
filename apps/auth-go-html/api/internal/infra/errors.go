package infra

import "net/http"

type Errors struct {
	logger *Logger
	json   *JSON
}

func NewErrors(l *Logger, j *JSON) *Errors {
	return &Errors{logger: l, json: j}
}

// errorResponse is a generic helper for sending JSON-formatted error
func (e *Errors) errorResponse(w http.ResponseWriter, r *http.Request, status int, errors map[string]string) {
	env := Envelope{"errors": errors}

	// Write the response using the writeJSON() helper. If this happens to return an
	// error then log it, and fall back to sending the client an empty response with a
	// 500 Internal Server Error status code.
	err := e.json.Write(w, status, env, nil)
	if err != nil {
		e.logError(r, err)
		w.WriteHeader(500)
	}
}

func (e *Errors) BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	e.errorResponse(w, r, http.StatusBadRequest, e.Error("bad_request", err.Error()))
}

// Errors has the type map[string]string, which is exactly
// the same as the errors map contained in the Validator type.
func (e *Errors) FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	e.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

// ServerErrorResponse will be used when our application encounters an
// unexpected problem at runtime. It logs the detailed error message, then uses
// errorResponse to send a 500 Internal Server Error status code and JSON
// response (containing a generic error message) to the client.
func (e *Errors) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	e.logError(r, err)

	message := "we got a problem on our side and could not process your request"
	e.errorResponse(w, r, http.StatusInternalServerError, e.Error("internal", message))
}

func (e *Errors) Error(code string, err string) map[string]string {
	return map[string]string{code: err}
}

// The logError() method is a generic helper for logging an error message.
func (e *Errors) logError(r *http.Request, err error) {
	e.logger.PrintError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}
