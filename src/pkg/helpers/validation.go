package helpers

import (
	"github.com/go-playground/locales/fa"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	fa_translations "github.com/go-playground/validator/v10/translations/fa"
)

//var translator ut.Translator
var validate *validator.Validate

func Translator() ut.Translator {

	farsi := fa.New()
	uni := ut.New(farsi, farsi)
	translator, _ := uni.GetTranslator("fa")
	return translator
}

func Validate(translator ut.Translator) *validator.Validate {
	if validate != nil {
		return validate
	}
	validate = validator.New()
	err := fa_translations.RegisterDefaultTranslations(validate, translator)
	if err != nil {
		panic(err)
		return nil
	}
	//overwriteTranslation(validate, Translator())

	return validate
}

func overwriteTranslation(validate *validator.Validate, translator ut.Translator) {

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
