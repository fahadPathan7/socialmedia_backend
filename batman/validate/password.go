package validate

import (
	"regexp"
	"strings"

	"github.com/fahadPathan7/socialmedia_backend/batman"
)

// Password validates password data type
func Password() func(dataI interface{}, fieldName string) []batman.ErrorDetails {
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

		pass := data
		// validate password length
		if len(pass) < 8 {
			errs = append(errs,
				batman.ErrorDetails{
					Field:   fieldName,
					Error:   INVALIDPASSWORD,
					Message: "Password needs to be at least 8 characters long",
				},
			)
		}

		// validate password upper case and lower case
		if strings.ToLower(pass) == pass || strings.ToUpper(pass) == pass {
			errs = append(errs,
				batman.ErrorDetails{
					Field:   fieldName,
					Error:   INVALIDPASSWORD,
					Message: "Password needs to include both lower and upper case",
				},
			)
		}

		// validate password number or special character
		rNum, _ := regexp.Compile(".*[0-9]")
		rSpecial, _ := regexp.Compile(".*[●!\"#$%&'()*+,\\-.\\/:;<=>?@\\[\\\\\\]^_`{|}~]")
		if !rNum.MatchString(pass) && !rSpecial.MatchString(pass) {
			errs = append(errs,
				batman.ErrorDetails{
					Field:   fieldName,
					Error:   INVALIDPASSWORD,
					Message: "Password needs to include at least 1 number or symbol",
				},
			)
		}

		return errs
	}
}
