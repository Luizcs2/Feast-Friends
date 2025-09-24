package helpers

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct performs validation on the given struct based on validate tags.
// It takes any struct as an interface{} parameter and validates its fields according to the defined validation rules.
// Returns an error if validation fails, or nil if the struct is valid.
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// FormatTime converts an RFC3339 formatted time string to a more readable format.
// This helper function is used to format timestamps in API models.
// Input: CreatedAt - string in RFC3339 format (e.g. "2006-01-02T15:04:05Z07:00")
// Output: formatted string in "DD MMM YYYY HH:mm" format (e.g. "02 Jan 2006 15:04")
// Returns empty string if parsing fails
func FormatTime(CreatedAt string) string {
	t, err := time.Parse(time.RFC3339, CreatedAt)
	if err != nil {
		return ""
	}
	return t.Format("02 Jan 2006 15:04")
}
