package apierr

import "net/http"

// Body of the error.
// Contains basic error info.
type Body struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

// APIError custom API error.
type APIError struct {
	Body         Body        `json:"error"`
	Meta         interface{} `json:"meta,omitempty"`
	HTTPCode     int         `json:"-"`
	Context      interface{} `json:"-"`
	ShouldReport bool        `json:"-"`
	HasTrace     bool        `json:"-"`
}

// Error returns error message.
func (e *APIError) Error() string {
	return e.Body.Message
}

// AddContext to the error.
// Only logged.
func (e *APIError) AddContext(context interface{}) *APIError {
	e.Context = context

	return e
}

// AddMeta to the error.
// Will be returned to the user.
func (e *APIError) AddMeta(meta interface{}) *APIError {
	e.Meta = meta

	return e
}

// WithTrace enables stacktrace reporting in logs.
func (e *APIError) WithTrace() *APIError {
	e.HasTrace = true

	return e
}

// Report error to the logs.
func (e *APIError) Report() *APIError {
	e.ShouldReport = true

	return e
}

// Send - helper for panic.
func (e *APIError) Send() {
	panic(e)
}

// WantsToBeReported in logs or not.
func (e *APIError) WantsToBeReported() bool {
	return e.ShouldReport
}

// WantsToShowTrace in logs or not.
func (e *APIError) WantsToShowTrace() bool {
	return e.HasTrace
}

// InternalServerError throw for 500.
var InternalServerError = &APIError{
	Body: Body{
		ID:      "internal_server_error",
		Message: "The server encountered an internal error or misconfiguration and was unable to complete your request.",
	},
	HTTPCode:     http.StatusInternalServerError,
	ShouldReport: true,
	HasTrace:     true,
}

// Forbidden throw for 403.
var Forbidden = &APIError{
	Body: Body{
		ID:      "forbidden",
		Message: "You don't have permissions to perform this request.",
	},
	HTTPCode: http.StatusForbidden,
}

// Unauthorized throw for 401.
var Unauthorized = &APIError{
	Body: Body{
		ID:      "invalid_credentials",
		Message: "Sent credentials are invalid.",
	},
	HTTPCode: http.StatusUnauthorized,
}

// NotFound throw for 404.
var NotFound = &APIError{
	Body: Body{
		ID:      "not_found",
		Message: "Requested object not found.",
	},
	HTTPCode: http.StatusNotFound,
}

// BadRequest throw for 402.
var BadRequest = &APIError{
	Body: Body{
		ID:      "bad_request",
		Message: "The server cannot process the request due to its malformed syntax.",
	},
	HTTPCode: http.StatusBadRequest,
}

// ValidationErrors is a set of messages for every invalid field.
type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

// ValidationError for single field.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValiationFailed throw for 422.
var ValiationFailed = &APIError{
	Body: Body{
		ID:      "validation_failed",
		Message: "Validation failed.",
	},
	HTTPCode: http.StatusUnprocessableEntity,
}
