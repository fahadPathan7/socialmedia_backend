package validate

import (
	"fmt"

	"github.com/fahadPathan7/socialmedia_backend/batman"
)

const (
	// MaxLimitIntLength defines the max value of limit parameter
	MaxLimitIntLength = 500
)

// Int64MaxLimit validate if data is non-nil
func Int64MaxLimit(maxLimit int64) func(dataI interface{}, fieldName string) []batman.ErrorDetails {
	return func(dataI interface{}, fieldName string) []batman.ErrorDetails {
		errs := []batman.ErrorDetails{}

		// check if data is valid string
		data, ok := dataI.(int64)
		if !ok {
			errs = append(errs, batman.ErrorDetails{
				Field:   fieldName,
				Error:   INVALID,
				Message: "Value is not valid",
			})
			return errs
		}

		if data > maxLimit {
			errs = append(errs, batman.ErrorDetails{
				Field:   fieldName,
				Error:   MAXLIMIT,
				Message: fmt.Sprintf("Maximum %d is allowed", maxLimit),
			})
		}

		return errs
	}
}
