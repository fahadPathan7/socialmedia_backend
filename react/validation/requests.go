package validation

import (
	"sync"

	// "github.com/fahadPathan7/socialmedia_backend/batman"
	// "github.com/fahadPathan7/socialmedia_backend/batman/validate"
	// pb "github.com/fahadPathan7/socialmedia_backend/proto/react"
)

var instantiated RequestValidator
var once sync.Once

type RequestValidator interface {
	// need to implement
}

type requestValidator struct{}


// NewRequestValidator returns a new RequestValidator object
func NewRequestValidator() *RequestValidator {
	once.Do(func() {
		instantiated = &requestValidator{}
	})
	return &instantiated
}