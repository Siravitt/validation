package validation

// built-in validator tags
const (
	TagRequired = "required"
	TagEmail    = "email"
	TagMin      = "min"
	TagMax      = "max"
	TagLen      = "len"
	TagNumeric  = "numeric"
	TagOneOf    = "oneof"
	TagURL      = "url"
	TagUUID4    = "uuid4"
	TagGT       = "gt"
	TagGTE      = "gte"
	TagLT       = "lt"
	TagLTE      = "lte"
)

// custom validator tags
const (
	TagThaiMobile = "thai_mobile"
	TagNoScript   = "no_script"
	TagNoEmoji    = "no_emoji"
)
