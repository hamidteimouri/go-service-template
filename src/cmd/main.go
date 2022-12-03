package main

import (
	"github.com/hamidteimouri/htutils/htapplife"
	"goservicetemplate/cmd/bootstrap"
)

func main() {
	htapplife.Start()
	bootstrap.Start()
	htapplife.ClosingSignalListener()
}
