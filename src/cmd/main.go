package main

import (
	"github.com/hamidteimouri/htutils/applife"
	"laramanpurego/cmd/bootstrap"
	"laramanpurego/cmd/di"
	"laramanpurego/internal/presentation/http"
)

func main() {
	applife.Start()
	bootstrap.Start()
	di.DB()
	http.StartHttp()

}
