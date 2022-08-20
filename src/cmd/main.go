package main

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"laramanpurego/src/internal/data/database"
	"laramanpurego/src/internal/presentation/http"
)

var (
	Validate           *validator.Validate
	UniversalTranslate ut.UniversalTranslator
)

func main() {
	database.Start()
	http.StartHttp()
}