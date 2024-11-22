package validation

import (
	"sync"

	"github.com/fahadPathan7/socialmedia_backend/batman"
	"github.com/fahadPathan7/socialmedia_backend/batman/validate"
	pb "github.com/fahadPathan7/socialmedia_backend/proto/user"
)

var instantiated RequestValidator
var once sync.Once

type RequestValidator interface {
	ValidateRegisterRequest(r *pb.RegisterRequest) error
	ValidateLoginRequest(r *pb.LoginRequest) error
}

type requestValidator struct{}

func (v *requestValidator) ValidateRegisterRequest(r *pb.RegisterRequest) error {
	errs := batman.Errors{}

	errs = validate.Field(
		r.GetUsername(), "username", errs,
		validate.StrRequired(),
		validate.StrMaxMin(validate.MaxStrLength, -1),
	)

	errs = validate.Field(
		r.GetPassword(), "password", errs,
		validate.StrRequired(),
		validate.Password(),
		validate.StrMaxMin(validate.MaxStrLength, -1),
	)

	errs = validate.Field(
		r.GetEmail(), "email", errs,
		validate.StrRequired(),
		validate.Email(),
		validate.StrMaxMin(validate.MaxStrLength, -1),
	)

	return errs
}

func (v *requestValidator) ValidateLoginRequest(r *pb.LoginRequest) error {
	errs := batman.Errors{}

	errs = validate.Field(
		r.GetEmail(), "email", errs,
		validate.StrRequired(),
		validate.Email(),
		validate.StrMaxMin(validate.MaxStrLength, -1),
	)

	errs = validate.Field(
		r.GetPassword(), "password", errs,
		validate.StrRequired(),
		validate.StrMaxMin(validate.MaxStrLength, -1),
	)

	return errs
}


// NewRequestValidator returns a new RequestValidator object
func NewRequestValidator() *RequestValidator {
	once.Do(func() {
		instantiated = &requestValidator{}
	})
	return &instantiated
}
