package validation

import (
	"sync"
	// "context"

	// "github.com/fahadPathan7/socialmedia_backend/batman"
	// "github.com/fahadPathan7/socialmedia_backend/batman/validate"
	// pb "github.com/fahadPathan7/socialmedia_backend/proto/post"
)

var instantiated RequestValidator
var once sync.Once

type RequestValidator interface {
	// ValidateCreateRequest(r *pb.CreateRequest) error
}

type requestValidator struct{}

// func (v *requestValidator) ValidateCreateRequest(r *pb.CreateRequest) error {
// 	// Validate the request
// 	err := validate.Validate(r)
// 	if err != nil {
// 		return batman.ErrValidationFailed
// 	}
// 	return nil
// }



// NewRequestValidator returns a new RequestValidator object
func NewRequestValidator() *RequestValidator {
	once.Do(func() {
		instantiated = &requestValidator{}
	})
	return &instantiated
}
