package validators

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var re = regexp.MustCompile(`fuck|nigger|cunt|microsoft`)

func ValidateTitleIsOk(field validator.FieldLevel) bool {
	checkStr := field.Field().String()
	checkStr = strings.ToLower(checkStr)

	return !re.MatchString(checkStr)
}
