package validation

import (
	"regexp"

	"github.com/forPelevin/gomoji"
	"github.com/go-playground/validator/v10"
)

var (
	reMobile = regexp.MustCompile(`^0[689]\d{8}$`)
	reScript = regexp.MustCompile(`(?i)<\s*/?\s*script\b`)
)

func registerCustomValidations() {
	validate.RegisterValidation(TagThaiMobile, validateThaiMobile)
	RegisterTranslation(TagThaiMobile, "{0} เบอร์โทรไม่ถูกต้อง")

	validate.RegisterValidation(TagNoScript, validateNoScript)
	RegisterTranslation(TagNoScript, "{0} ห้ามใส่ <script>")

	validate.RegisterValidation(TagNoEmoji, validateNoEmoji)
	RegisterTranslation(TagNoEmoji, "{0} ห้ามใส่ emoji")
}

func validateThaiMobile(fl validator.FieldLevel) bool {
	return reMobile.MatchString(fl.Field().String())
}

func validateNoScript(fl validator.FieldLevel) bool {
	return !reScript.MatchString(fl.Field().String())
}

func validateNoEmoji(fl validator.FieldLevel) bool {
	return !gomoji.ContainsEmoji(fl.Field().String())
}
