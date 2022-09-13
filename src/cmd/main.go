package main

import (
	"github.com/hamidteimouri/htutils/htapplife"
	"laramanpurego/cmd/bootstrap"
)

func main() {
	htapplife.Start()
	bootstrap.Start()
	htapplife.ClosingSignalListener()

}
