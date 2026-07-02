package validation

func overrideThaiTranslations() {
	RegisterTranslation(TagRequired, "{0} ไม่ได้ระบุค่า")
	RegisterTranslation(TagEmail, "{0} รูปแบบอีเมลไม่ถูกต้อง")
	RegisterTranslation(TagMin, "{0} ต้องมีความยาวอย่างน้อย {1}")
	RegisterTranslation(TagMax, "{0} ต้องมีความยาวไม่เกิน {1}")
	RegisterTranslation(TagLen, "{0} ต้องมีความยาว {1}")
	RegisterTranslation(TagNumeric, "{0} ต้องเป็นตัวเลขเท่านั้น")
	RegisterTranslation(TagOneOf, "{0} ต้องเป็นค่าใดค่าหนึ่งใน {1}")
	RegisterTranslation(TagURL, "{0} รูปแบบ URL ไม่ถูกต้อง")
	RegisterTranslation(TagUUID4, "{0} รูปแบบ UUID ไม่ถูกต้อง")
	RegisterTranslation(TagGT, "{0} ต้องมีค่ามากกว่า {1}")
	RegisterTranslation(TagGTE, "{0} ต้องมีค่ามากกว่าหรือเท่ากับ {1}")
	RegisterTranslation(TagLT, "{0} ต้องมีค่าน้อยกว่า {1}")
	RegisterTranslation(TagLTE, "{0} ต้องมีค่าน้อยกว่าหรือเท่ากับ {1}")
}
