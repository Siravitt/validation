package validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- Validate (struct) ---

func TestValidate_Pass(t *testing.T) {
	type Req struct {
		Name  string `json:"name"  validate:"required,max=100"`
		Email string `json:"email" validate:"required,email"`
	}
	err := Validate(Req{Name: "John", Email: "john@example.com"})
	assert.NoError(t, err)
}

func TestValidate_Fail_ReturnsValidationErrors(t *testing.T) {
	type Req struct {
		Name  string `json:"name"  validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}
	err := Validate(Req{})
	require.Error(t, err)

	errs, ok := err.(ValidationErrors)
	require.True(t, ok)
	assert.Len(t, errs, 2)
}

func TestValidate_FieldName_UsesJsonTag(t *testing.T) {
	type Req struct {
		FirstName string `json:"first_name" validate:"required"`
	}
	err := Validate(Req{})
	require.Error(t, err)

	errs := err.(ValidationErrors)
	assert.Equal(t, "first_name", errs[0].Field)
}

func TestValidate_DefaultThaiMessage_Required(t *testing.T) {
	type Req struct {
		Name string `json:"name" validate:"required"`
	}
	err := Validate(Req{})
	require.Error(t, err)

	errs := err.(ValidationErrors)
	assert.Equal(t, "name ไม่ได้ระบุค่า", errs[0].Message)
}

func TestValidate_DefaultThaiMessage_Email(t *testing.T) {
	type Req struct {
		Email string `json:"email" validate:"required,email"`
	}
	err := Validate(Req{Email: "not-an-email"})
	require.Error(t, err)

	errs := err.(ValidationErrors)
	assert.Equal(t, "email รูปแบบอีเมลไม่ถูกต้อง", errs[0].Message)
}

func TestValidate_DefaultThaiMessage_MinMax(t *testing.T) {
	type Req struct {
		Name string `json:"name" validate:"min=5"`
		Bio  string `json:"bio"  validate:"max=3"`
	}
	err := Validate(Req{Name: "ab", Bio: "toolong"})
	require.Error(t, err)

	errs := err.(ValidationErrors)
	assert.Equal(t, "name ต้องมีความยาวอย่างน้อย 5", errs[0].Message)
	assert.Equal(t, "bio ต้องมีความยาวไม่เกิน 3", errs[1].Message)
}

// --- ValidateVar ---

func TestValidateVar_Pass(t *testing.T) {
	err := ValidateVar("john@example.com", "required", "email")
	assert.NoError(t, err)
}

func TestValidateVar_Fail(t *testing.T) {
	err := ValidateVar("", "required")
	assert.Error(t, err)
}

func TestValidateVar_MultiTag(t *testing.T) {
	err := ValidateVar("ab", "required", "min=5")
	assert.Error(t, err)
}

// --- ValidateWithKey ---

func TestValidateWithKey_Pass(t *testing.T) {
	err := ValidateWithKey("user_id", "abc-123", "required")
	assert.NoError(t, err)
}

func TestValidateWithKey_Fail_ContainsKey(t *testing.T) {
	err := ValidateWithKey("user_id", "", "required")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "user_id")
	assert.Contains(t, err.Error(), "required")
}

// --- RegisterTranslation ---

func TestRegisterTranslation_OverridesMessage(t *testing.T) {
	RegisterTranslation("required", "{0} ห้ามว่าง (custom)")

	type Req struct {
		Title string `json:"title" validate:"required"`
	}
	err := Validate(Req{})
	require.Error(t, err)

	errs := err.(ValidationErrors)
	assert.Equal(t, "title ห้ามว่าง (custom)", errs[0].Message)

	// restore default
	RegisterTranslation("required", "{0} ไม่ได้ระบุค่า")
}

// --- RegisterValidation ---

func TestRegisterValidation_CustomTag(t *testing.T) {
	err := RegisterValidation("must_hello", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "hello"
	}, "{0} ต้องเป็น hello เท่านั้น")
	require.NoError(t, err)

	type Req struct {
		Word string `json:"word" validate:"must_hello"`
	}

	assert.NoError(t, Validate(Req{Word: "hello"}))

	err = Validate(Req{Word: "world"})
	require.Error(t, err)
	errs := err.(ValidationErrors)
	assert.Equal(t, "word ต้องเป็น hello เท่านั้น", errs[0].Message)
}

// --- Custom validators ---

func TestThaiMobile_Valid(t *testing.T) {
	cases := []string{"0812345678", "0912345678", "0612345678"}
	for _, c := range cases {
		assert.NoError(t, ValidateVar(c, "thai_mobile"), c)
	}
}

func TestThaiMobile_Invalid(t *testing.T) {
	cases := []string{"", "1234567890", "08123456", "0712345678"}
	for _, c := range cases {
		assert.Error(t, ValidateVar(c, "thai_mobile"), c)
	}
}

func TestNoScript_Valid(t *testing.T) {
	cases := []string{"hello world", "<b>bold</b>", "normal text"}
	for _, c := range cases {
		assert.NoError(t, ValidateVar(c, "no_script"), c)
	}
}

func TestNoScript_Invalid(t *testing.T) {
	cases := []string{"<script>alert(1)</script>", "<SCRIPT>", "</script>"}
	for _, c := range cases {
		assert.Error(t, ValidateVar(c, "no_script"), c)
	}
}

func TestNoEmoji_Valid(t *testing.T) {
	cases := []string{"hello world", "สวัสดี", "normal text 123"}
	for _, c := range cases {
		assert.NoError(t, ValidateVar(c, TagNoEmoji), c)
	}
}

func TestNoEmoji_Invalid(t *testing.T) {
	cases := []string{"hello 😀", "🎉 party", "text 🔥 fire"}
	for _, c := range cases {
		assert.Error(t, ValidateVar(c, TagNoEmoji), c)
	}
}

// --- ValidationErrors ---

func TestValidationErrors_ErrorIsJSON(t *testing.T) {
	type Req struct {
		Name string `json:"name" validate:"required"`
	}
	err := Validate(Req{})
	require.Error(t, err)

	assert.Contains(t, err.Error(), `"field"`)
	assert.Contains(t, err.Error(), `"message"`)
}

func TestNoEmoji_Message(t *testing.T) {
	type Req struct {
		Name string `json:"name" validate:"no_emoji"`
	}
	err := Validate(Req{Name: "hello 😀"})
	require.Error(t, err)

	errs := err.(ValidationErrors)
	assert.Equal(t, "name ห้ามใส่ emoji", errs[0].Message)
}

// --- extractErrors: InvalidValidationError branch ---

func TestValidate_NonStruct_ReturnsError(t *testing.T) {
	// passing a non-struct triggers *validator.InvalidValidationError
	err := Validate("not a struct")
	assert.Error(t, err)
}

// --- RegisterValidation: error branch ---

func TestRegisterValidation_NilFunc_ReturnsError(t *testing.T) {
	err := RegisterValidation("nil_tag", nil, "msg")
	assert.Error(t, err)
}
