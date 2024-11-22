package validate

import "github.com/fahadPathan7/socialmedia_backend/batman"

// Validator interface def
type Validator interface {
	Validate(data interface{}) *batman.ErrorDetails
}

// Field validates a field
func Field(data interface{}, fieldName string, errs batman.Errors, validators ...func(data interface{}, fieldName string) []batman.ErrorDetails) batman.Errors {
	// errs := batman.Errors{}

	for _, validator := range validators {
		errDetailsArray := validator(data, fieldName)
		if len(errDetailsArray) > 0 {
			errs.Errors = append(errs.Errors, errDetailsArray...)
		}
	}

	return errs
}
