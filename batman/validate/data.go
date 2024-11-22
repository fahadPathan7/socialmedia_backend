package validate

import "github.com/fahadPathan7/socialmedia_backend/batman"

const (
	// MaxArrayLength defines the max size of array it can process at a time
	MaxArrayLength = 100
)

// NotNil validate if data is non-nil
func NotNil() func(dataI interface{}, fieldName string) []batman.ErrorDetails {
	return func(dataI interface{}, fieldName string) []batman.ErrorDetails {
		errs := []batman.ErrorDetails{}
		if dataI == nil {
			errs = append(errs, batman.ErrorDetails{
				Field:   fieldName,
				Error:   REQUIRED,
				Message: "Value is required",
			})
		}

		return errs
	}
}
