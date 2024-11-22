package validate

import (
	"fmt"

	"github.com/fahadPathan7/socialmedia_backend/batman"
)

const (
	// MaxStrLength defines the max string length from input
	MaxStrLength = 255
	// MaxStrDescriptionLength is the max length of larger input string
	MaxStrDescriptionLength = 2000
)

// StrRequired validate if data is non-empty
func StrRequired() func(dataI interface{}, fieldName string) []batman.ErrorDetails {
	return func(dataI interface{}, fieldName string) []batman.ErrorDetails {
		errs := []batman.ErrorDetails{}

		// check if data is valid string
		data, ok := dataI.(string)
		if !ok {
			errs = append(errs, batman.ErrorDetails{
				Field:   fieldName,
				Error:   INVALID,
				Message: "Value is not valid",
			})
			return errs
		}

		if len(data) == 0 {
			errs = append(errs, batman.ErrorDetails{
				Field:   fieldName,
				Error:   REQUIRED,
				Message: "Value is required",
			})
		}

		return errs
	}
}

// StrMaxMin validates max and min length of a data
func StrMaxMin(max, min int) func(dataI interface{}, fieldName string) []batman.ErrorDetails {
	return func(dataI interface{}, fieldName string) []batman.ErrorDetails {
		errs := []batman.ErrorDetails{}

		// check if data is valid string
		data, ok := dataI.(string)
		if !ok {
			errs = append(errs, batman.ErrorDetails{
				Field:   fieldName,
				Error:   INVALID,
				Message: "Value is not valid",
			})
			return errs
		}

		if len(data) > max && max != -1 {
			errs = append(errs, batman.ErrorDetails{
				Field:   fieldName,
				Error:   MAXLIMIT,
				Message: fmt.Sprintf("Maximum %d character limit exceeded", max),
			})
		}

		if len(data) < min && min != -1 {
			errs = append(errs, batman.ErrorDetails{
				Field:   fieldName,
				Error:   MINLIMIT,
				Message: fmt.Sprintf("Minimum %d characters required", min),
			})
		}
		return errs
	}
}
