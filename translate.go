package validation

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func registerDefaultTranslations() {
	override := func(tag, msg string) {
		validate.RegisterTranslation(tag, trans,
			func(ut ut.Translator) error { return ut.Add(tag, msg, true) },
			func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T(tag, fe.Field(), fe.Param())
				return t
			},
		)
	}

	override("required", "{0} ไม่ได้ระบุค่า")
	override("email", "{0} รูปแบบอีเมลไม่ถูกต้อง")
	override("min", "{0} ต้องมีความยาวอย่างน้อย {1}")
	override("max", "{0} ต้องมีความยาวไม่เกิน {1}")
	override("len", "{0} ต้องมีความยาว {1}")
	override("numeric", "{0} ต้องเป็นตัวเลขเท่านั้น")
	override("oneof", "{0} ต้องเป็นค่าใดค่าหนึ่งใน {1}")
	override("url", "{0} รูปแบบ URL ไม่ถูกต้อง")
	override("uuid4", "{0} รูปแบบ UUID ไม่ถูกต้อง")
	override("gt", "{0} ต้องมีค่ามากกว่า {1}")
	override("gte", "{0} ต้องมีค่ามากกว่าหรือเท่ากับ {1}")
	override("lt", "{0} ต้องมีค่าน้อยกว่า {1}")
	override("lte", "{0} ต้องมีค่าน้อยกว่าหรือเท่ากับ {1}")
}
