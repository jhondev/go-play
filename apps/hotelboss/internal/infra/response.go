package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Response struct {
	*Logger
	*Errors
}
type Envelope map[string]any

func NewResponse(l *Logger, e *Errors) *Response {
	return &Response{Logger: l, Errors: e}
}

func (_ *Response) Write(w http.ResponseWriter, status int, data Envelope, headers http.Header) error {
	// js, err := json.Marshal(data)
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (r *Response) Success(w http.ResponseWriter, data Envelope) error {
	return r.Write(w, http.StatusOK, data, nil)
}

func (r *Response) Created(w http.ResponseWriter, data Envelope) error {
	return r.Write(w, http.StatusCreated, data, nil)
}

func (r *Response) BadRequest(w http.ResponseWriter, req *http.Request, err error) {
	r.error(w, req, http.StatusBadRequest, r.Error("bad_request", err.Error()))
}

// ServerError will be used when our application encounters an
// unexpected problem at runtime. It logs the detailed error message, then uses
// errorResponse to send a 500 Internal Server Error status code and JSON
// response (containing a generic error message) to the client.
func (r *Response) ServerError(w http.ResponseWriter, req *http.Request, err error) {
	r.LogError(req, err)

	message := "we got a problem on our side and could not process your request"
	r.error(w, req, http.StatusInternalServerError, r.Error("internal", message))
}

// errorResponse is a generic helper for sending JSON-formatted error
func (r *Response) error(w http.ResponseWriter, req *http.Request, status int, errors map[string]string) {
	env := Envelope{"errors": errors}

	// Write the response using the write() helper. If this happens to return an
	// error then log it, and fall back to sending the client an empty response with a
	// 500 Internal Server Error status code.
	err := r.Write(w, status, env, nil)
	if err != nil {
		r.LogError(req, err)
		w.WriteHeader(500)
	}
}

func (_ *Response) Read(w http.ResponseWriter, r *http.Request, dst any) error {
	// Use http.MaxBytesReader() to limit the size of the request body to 1MB.
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// Initialize the json.Decoder, and call the DisallowUnknownFields() method on it
	// before decoding. This means that if the JSON from the client now includes any
	// field which cannot be mapped to the target destination, the decoder will return
	// an error instead of just ignoring the field.
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	// Decode the request body to the destination.
	err := dec.Decode(dst)

	if err != nil {
		// If there is an error during decoding, start the triage...
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		// Use the errors.As() function to check whether the error has the type
		// *json.SyntaxError. If it does, then return a plain-english error message
		// which includes the location of the problem.
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		// In some circumstances Decode() may also return an io.ErrUnexpectedEOF error
		// for syntax errors in the JSON. So we check for this using errors.Is() and
		// return a generic error message. There is an issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		// Likewise, catch any *json.UnmarshalTypeError errors. These occur when the
		// JSON value is the wrong type for the target destination. If the error relates
		// to a specific field, then we include that in our error message to make it
		// easier for the client to debug.
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		// An io.EOF error will be returned by Decode() if the request body is empty. We
		// check for this with errors.Is() and return a plain-english error message
		// instead.
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		// If the JSON contains a field which cannot be mapped to the target destination
		// then Decode() will now return an error message in the format "json: unknown
		// field "<name>"". We check for this, extract the field name from the error,
		// and interpolate it into our custom error message. Note that there's an open
		// issue at https://github.com/golang/go/issues/29035 regarding turning this
		// into a distinct error type in the future.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		// If the request body exceeds 1MB in size the decode will now fail with the
		// error "http: request body too large". There is an open issue about turning
		// this into a distinct error type at https://github.com/golang/go/issues/30715.
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)

		// A json.InvalidUnmarshalError error will be returned if we pass something
		// that is not a non-nil pointer to Decode(). We catch this and panic,
		// rather than returning an error to our handler.
		case errors.As(err, &invalidUnmarshalError):
			panic(err)

			// For anything else, return the error message as-is.
		default:
			return err
		}
	}

	// Call Decode() again, using a pointer to an empty anonymous struct as the
	// destination. If the request body only contained a single JSON value this will
	// return an io.EOF error. So if we get anything else, we know that there is
	// additional data in the request body and we return our own custom error message.
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}
