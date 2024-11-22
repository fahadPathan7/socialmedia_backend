package validate

import (
	"regexp"

	"github.com/fahadPathan7/socialmedia_backend/batman"
)

// Email validates email data type
func Email() func(dataI interface{}, fieldName string) []batman.ErrorDetails {
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

		// Check if email is valid
		rEmail := regexp.MustCompile("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$")
		if !rEmail.MatchString(data) {
			errs = append(errs, batman.ErrorDetails{
				Field:   fieldName,
				Error:   INVALIDEMAIL,
				Message: "Email is not valid",
			})
		}

		return errs
	}
}
