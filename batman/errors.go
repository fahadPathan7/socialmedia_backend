package batman

import (
	"errors"
	"strings"

	"github.com/fahadPathan7/socialmedia_backend/proto/batman"
	bpb "github.com/fahadPathan7/socialmedia_backend/proto/batman"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidationError holds all validation errors
type ValidationError struct {
	Errors []string
}

// Error struct
type Error struct {
	Field   string
	Message string
	Error   string
}

// Errors struct
type Errors struct {
	Errors []ErrorDetails
}

// ErrorDetails data type
type ErrorDetails = bpb.ErrorDetails

const (
	// EmailExists an email
	EmailExists = "email already exists"
	// EmailNotSent an email
	EmailNotSent = "could not send email"
	// ValidationFailed message
	ValidationFailed = "validation failed"
	// InvalidData message
	InvalidData = "invalid data"
	// InvalidPayload message
	InvalidPayload = "invalid payload"
	// NotFound message
	NotFound = "not found"
	// CouldNotFetch message
	CouldNotFetch = "could not fetch data"
	// CarNotAvailable error
	CarNotAvailable = "car not available"
)

// errors

// ErrEmailNotSent an email
var ErrEmailNotSent = errors.New("could not send email")

// ErrValidationFailed error
var ErrValidationFailed = errors.New("validation failed")

// ErrInvalidData error
var ErrInvalidData = errors.New("invalid data")

// ErrNotFound error
var ErrNotFound = errors.New("not found")

// ErrUnknown error
var ErrUnknown = errors.New("something went wrong")

// ErrCouldNotFetch err
var ErrCouldNotFetch = errors.New("could not fetch")

// ErrCouldNotFetch err
var ErrCouldNotDelete = errors.New("could not delete")

// ComposeMultipleErrorStr composes multiple error strings to a returnable status.Status
func ComposeMultipleErrorStr(code codes.Code, message string, errs []string) *status.Status {
	st := status.New(code, message)

	for _, er := range errs {
		st, _ = st.WithDetails(&batman.ErrorDetails{
			Error:   er,
			Message: er,
		})
	}
	return st
}

// ComposeErrors composes multiple error details from Errors to a returnable status.Status
func ComposeErrors(code codes.Code, message string, errs Errors) *status.Status {
	st := status.New(code, message)

	for _, err := range errs.Errors {
		st, _ = st.WithDetails(&batman.ErrorDetails{
			Error:   err.Error,
			Message: err.Message,
			Field:   err.Field,
		})
	}

	return st
}

// Implement the Error method for Errors type. This will allow Errors to be returned as an error type
func (e Errors) Error() string {
    var errorMessages []string
    for _, err := range e.Errors {
        errorMessages = append(errorMessages, err.Message)
    }
    return strings.Join(errorMessages, ", ")
}