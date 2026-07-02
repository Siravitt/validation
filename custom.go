package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	reMobile = regexp.MustCompile(`^0[689]\d{8}$`)
	reScript = regexp.MustCompile(`(?i)<\s*/?\s*script\b`)
)

func registerCustomValidations() {
	validate.RegisterValidation("thai_mobile", validateThaiMobile)
	RegisterTranslation("thai_mobile", "{0} เบอร์โทรไม่ถูกต้อง")

	validate.RegisterValidation("no_script", validateNoScript)
	RegisterTranslation("no_script", "{0} ห้ามใส่ <script>")
}

func validateThaiMobile(fl validator.FieldLevel) bool {
	return reMobile.MatchString(fl.Field().String())
}

func validateNoScript(fl validator.FieldLevel) bool {
	return !reScript.MatchString(fl.Field().String())
}
