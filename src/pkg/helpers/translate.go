package helpers

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TranslateValidation(validate *validator.Validate, translator ut.Translator) {

	/* Overwrite Email */
	validate.RegisterTranslation("email", translator, func(ut ut.Translator) error {
		return ut.Add("email", "فیلد {0} معتبر نیست", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())

		return t
	})

	/* Overwrite Required */
	validate.RegisterTranslation("required", translator, func(ut ut.Translator) error {
		return ut.Add("required", "فیلد {0} الزامی است", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

}
