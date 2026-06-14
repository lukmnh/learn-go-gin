package validation

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func PasswordValidator(fieldLevel validator.FieldLevel) bool {
	p := fieldLevel.Field().String()

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(p)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(p)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(p)

	return hasUpper && hasLower && hasNumber
}
func RegisterValidation() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		return v.RegisterValidation("password", PasswordValidator)
	}
	return nil
}
