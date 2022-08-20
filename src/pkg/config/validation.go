package config

import (
	"github.com/go-playground/locales/fa"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	fa_translations "github.com/go-playground/validator/v10/translations/fa"
	"laramanpurego/pkg/helpers"
)

var translator ut.Translator

func Translator() ut.Translator {
	fa := fa.New()
	uni := ut.New(fa, fa)
	translator, _ = uni.GetTranslator("fa")
	return translator
}

func Validate() *validator.Validate {
	validate := validator.New()
	fa_translations.RegisterDefaultTranslations(validate, Translator())
	helpers.TranslateValidation(validate, Translator())

	return validate
}
