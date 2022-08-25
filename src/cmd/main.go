package main

import (
	"github.com/hamidteimouri/htutils/applife"
	"laramanpurego/cmd/di"
	"laramanpurego/internal/presentation/http"
)

func main() {
	applife.Start()
	di.DB()
	http.StartHttp()
}
