package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	t, _ := uni.GetTranslator("en")
	trans = t

	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" || name == "" {
			return fld.Name
		}
		return name
	})

	en_translations.RegisterDefaultTranslations(validate, trans)
	registerDefaultTranslations()
	registerCustomValidations()
}

// Validate validates a struct. Returns nil or ValidationErrors.
func Validate(s any) error {
	if err := validate.Struct(s); err != nil {
		return extractErrors(err)
	}
	return nil
}

// ValidateVar validates a single value against one or more tags.
func ValidateVar(field any, tags ...string) error {
	return validate.Var(field, strings.Join(tags, ","))
}

// ValidateWithKey validates a single value and wraps the error with a key name.
func ValidateWithKey(key string, field any, tags ...string) error {
	if err := ValidateVar(field, tags...); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return fmt.Errorf("'%s' failed on '%s' tag", key, e.ActualTag())
		}
	}
	return nil
}

// RegisterTranslation overrides the error message for a validation tag.
// Supports {0} = field name, {1} = param.
func RegisterTranslation(tag, msg string) {
	validate.RegisterTranslation(tag, trans,
		func(ut ut.Translator) error { return ut.Add(tag, msg, true) },
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(tag, fe.Field(), fe.Param())
			return t
		},
	)
}

// RegisterValidation adds a custom validation tag and its translator.
func RegisterValidation(tag string, fn validator.Func, msg string) error {
	if err := validate.RegisterValidation(tag, fn); err != nil {
		return err
	}
	RegisterTranslation(tag, msg)
	return nil
}
